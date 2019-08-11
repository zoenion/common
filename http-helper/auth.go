package http_helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"github.com/zoenion/common/errors"
	authpb "github.com/zoenion/common/proto/auth"
	"log"
	"net/http"
	"strings"
)

// CredentialsProvider
type CredentialsProvider func(args ...string) *authpb.Credentials

type JwtVerifier func(string) bool

// DigestAuthenticationMiddleware
type DigestAuthenticationMiddleware struct {
	digest   *auth.DigestAuth
	realm    string
	wrappers []RequestWrapper
}

func (dam *DigestAuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if username, _ := dam.digest.CheckAuth(r); username == "" {
			dam.digest.RequireAuth(w, r)
		} else {
			for _, wrapper := range dam.wrappers {
				r = wrapper(r)
			}
			next(w, r)
		}
	}
}

func NewDigestAuthenticationMiddleware(realm string, provider CredentialsProvider, wrappers ...RequestWrapper) *DigestAuthenticationMiddleware {
	a := &DigestAuthenticationMiddleware{
		realm:    realm,
		wrappers: wrappers,
	}
	a.digest = auth.NewDigestAuthenticator(realm, func(username, realm string) string {
		credentials := provider(username)
		if credentials == nil {
			return ""
		}

		ha1Str := fmt.Sprintf("%s:%s:%s", username, realm, credentials.Password)
		h := md5.New()
		h.Write([]byte(ha1Str))
		ha1Bytes := h.Sum(nil)
		return hex.EncodeToString(ha1Bytes)
	})
	return a
}

// BasicAuthenticationMiddleware
type BasicAuthenticationMiddleware struct {
	provider CredentialsProvider
	wrappers []RequestWrapper
	realm    string
}

func (bam *BasicAuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	reqAuth := &RequireAuth{
		Realm: bam.realm,
		Type:  "Basic",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if !ok {
			WriteResponse(w, 401, reqAuth)
			return
		}

		c := bam.provider(user)
		if c == nil || (c.Username != user && c.Email != user) || c.Password != password {
			WriteResponse(w, 401, reqAuth)
			return
		}

		for _, wrapper := range bam.wrappers {
			r = wrapper(r)
		}
		next.ServeHTTP(w, r)
	}
}

func NewBasicAuthenticationMiddleware(realm string, provider CredentialsProvider, wrappers ...RequestWrapper) *BasicAuthenticationMiddleware {
	return &BasicAuthenticationMiddleware{
		realm:    realm,
		provider: provider,
		wrappers: wrappers,
	}
}

// BearerAuthenticationMiddleware
type BearerAuthenticationMiddleware struct {
	jwtVerifier JwtVerifier
	wrappers    []RequestWrapper
}

func (bam *BearerAuthenticationMiddleware) Wrap(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if bam.jwtVerifier(authorization) {
			for _, wrapper := range bam.wrappers {
				r = wrapper(r)
			}
			next(w, r)
			return
		}
		WriteError(w, errors.HttpForbidden)
	}
}

func NewBearerAuthenticationMiddleware(jwtVerifier JwtVerifier, wrappers ...RequestWrapper) *BearerAuthenticationMiddleware {
	return &BearerAuthenticationMiddleware{
		jwtVerifier: jwtVerifier,
		wrappers:    wrappers,
	}
}

// ProxyAuthentication
type ProxyAuthenticationMiddleware struct {
	realm       string
	credentials *authpb.Credentials
	wrappers    []RequestWrapper
}

func (pam *ProxyAuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Proxy-Authentication")
		if authorizationHeader == "" {
			log.Println("Proxy authentication failed")
			w.WriteHeader(http.StatusProxyAuthRequired)
			w.Header().Set("Proxy-Authenticate", fmt.Sprintf("Basic realm=%s", pam.realm))
			return
		}

		authentication := strings.TrimLeft(authorizationHeader, "Basic ")
		data, err := base64.StdEncoding.DecodeString(authentication)
		if err != nil {
			log.Println("Proxy authentication failed")
			w.WriteHeader(http.StatusProxyAuthRequired)
			w.Header().Set("Proxy-Authenticate", fmt.Sprintf("Basic realm=%s", pam.realm))
			return
		}

		parts := strings.Split(string(data), ":")
		if len(parts) != 2 || pam.credentials.Username != parts[0] || pam.credentials.Password != parts[1] {
			w.WriteHeader(http.StatusProxyAuthRequired)
			w.Header().Set("Proxy-Authenticate", fmt.Sprintf("Basic realm=%s", pam.realm))
			return
		}

		for _, wrapper := range pam.wrappers {
			r = wrapper(r)
		}
		next(w, r)
	}
}

func NewProxyAuthenticationMiddleware(realm string, credentials *authpb.Credentials, wrappers ...RequestWrapper) *ProxyAuthenticationMiddleware {
	return &ProxyAuthenticationMiddleware{
		realm:       realm,
		credentials: credentials,
		wrappers:    wrappers,
	}
}
