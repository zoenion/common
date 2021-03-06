package grpcx

import (
	"context"
	"crypto"
	"crypto/x509"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/omecodes/common/errors"
	"github.com/omecodes/common/netx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"net"
	"net/http"
	"strings"
)

type Server struct {
	host         string
	options      options
	certificate  *x509.Certificate
	key          crypto.PrivateKey
	initialized  bool
	mux          *runtime.ServeMux
	grpcServer   *grpc.Server
	errorChannel chan error
	grpcAddress  string
	grpcListener net.Listener
	httpAddress  string
	httpListener net.Listener
	stopped      bool
}

func (s *Server) listenHttp() error {
	if s.options.endpointMappers != nil {

		address := fmt.Sprintf("%s:", s.host)
		if s.options.httpPort > 0 {
			address = fmt.Sprintf("%s%d", address, s.options.httpPort)
		}
		l, err := netx.Listen(address, s.options.listenOptions...)
		if err != nil {
			return err
		}
		s.httpListener = l

		if s.options.httpPort > 0 {
			s.httpAddress = address
		} else {
			s.httpAddress = address + strings.Split(l.Addr().String(), ":")[1]
		}

		//log.Info("[gRPC-http] starting HTTP server", log.Field("at", s.httpAddress))
	}
	return nil
}

func (s *Server) listenGRPC() error {

	address := fmt.Sprintf("%s:", s.host)
	if s.options.gRPCPort > 0 {
		address = fmt.Sprintf("%s%d", address, s.options.gRPCPort)
	}
	l, err := netx.Listen(address, s.options.listenOptions...)
	if err != nil {
		return err
	}
	s.grpcListener = l

	if s.options.gRPCPort > 0 {
		s.grpcAddress = address
	} else {
		s.grpcAddress = address + strings.Split(l.Addr().String(), ":")[1]
	}
	//log.Info("[gRPC-http] starting gRPC server", log.Field("at", s.grpcAddress))

	return nil
}

func (s *Server) init() error {
	if s.initialized {
		return nil
	}
	s.initialized = true
	s.errorChannel = make(chan error, 4)

	err := s.listenGRPC()
	if err != nil {
		return err
	}

	err = s.listenHttp()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) startGRPC() {
	if !s.initialized {
		s.handleError(errors.New("Init method must me called at least once"))
		return
	}

	server := s.GRPCServer()
	err := server.Serve(s.grpcListener)
	if err != nil {
		if !s.stopped {
			s.handleError(err)
		}
	}
}

func (s *Server) startHTTP() {

	if s.options.endpointMappers != nil {
		if !s.initialized {
			s.handleError(errors.New("Init method must me called at least once"))
			return
		}

		var serverOpts []runtime.ServeMuxOption
		if s.options.grpcSession {
			serverOpts = append(serverOpts, runtime.WithForwardResponseOption(SetCookieFromGRPCMetadata))
		}
		serverOpts = append(serverOpts, runtime.WithProtoErrorHandler(s.HandlerError))
		s.mux = runtime.NewServeMux(serverOpts...)

		var opts []grpc.DialOption

		var lopts netx.ListenOptions
		for _, o := range s.options.listenOptions {
			o(&lopts)
		}

		if lopts.Secure {
			if lopts.TLS != nil {
				cert := lopts.TLS.Certificates[0]
				c := credentials.NewServerTLSFromCert(&cert)
				opts = append(opts, grpc.WithTransportCredentials(c))
			} else {
				c, err := credentials.NewClientTLSFromFile(lopts.CertFilename, "")
				if err != nil {
					s.handleError(err)
					return
				}
				opts = append(opts, grpc.WithTransportCredentials(c))
			}
		} else {
			opts = append(opts, grpc.WithInsecure())
		}

		for _, m := range s.options.endpointMappers {
			err := m.Mapper(context.Background(), s.mux, s.grpcAddress, opts)
			if err != nil {
				s.handleError(err)
				return
			}
		}

		var handler http.Handler
		if len(s.options.middlewareList) > 0 {
			handler = s.mux
			for _, middleware := range s.options.middlewareList {
				handler = middleware(handler)
			}
		} else {
			handler = s.mux
		}

		err := http.Serve(s.httpListener, handler)
		if err != nil {
			if !s.stopped {
				s.handleError(err)
			}
		}
	}
}

func (s *Server) Start() error {
	err := s.init()
	if err != nil {
		return err
	}
	go s.startGRPC()
	go s.startHTTP()
	return nil
}

func (s *Server) Certificate() *x509.Certificate {
	return s.certificate
}

func (s *Server) Key() crypto.PrivateKey {
	return s.key
}

func (s *Server) GRPCServer() *grpc.Server {
	if s.options.authFunc == nil {
		s.options.authFunc = func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		}
	}
	if s.grpcServer == nil {
		s.grpcServer = grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_auth.StreamServerInterceptor(s.options.authFunc),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_auth.UnaryServerInterceptor(s.options.authFunc),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)
	}
	return s.grpcServer
}

func (s *Server) Stop() {
	s.stopped = true

	if s.httpListener != nil {
		_ = s.httpListener.Close()
	}

	if s.grpcListener != nil {
		_ = s.httpListener.Close()
	}
}

func (s *Server) GatewayAddress() string {
	return s.httpAddress
}

func (s *Server) GRPCAddress() string {
	return s.grpcAddress
}

func (s *Server) handleError(err error) {
	c := s.Errors()
	c <- err
}

func (s *Server) Errors() chan error {
	if s.errorChannel == nil {
		s.errorChannel = make(chan error, 4)
	}
	return s.errorChannel
}

func (s *Server) HandlerError(ctx context.Context, mux *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	//log.Info("caught error", log.Field("err", err))
	st, ok := status.FromError(err)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(errors.HttpStatus(errors.New(st.Message())))
}

func New(host string, opts ...Option) *Server {
	s := &Server{
		host:         host,
		errorChannel: make(chan error, 2),
	}

	for _, o := range opts {
		o(&s.options)
	}

	return s
}
