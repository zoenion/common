package backend

import (
	"encoding/base64"
	"github.com/omecodes/common/grpcx"
	"github.com/omecodes/common/utils/log"
	"net/http"
	"strings"
)

func ProxyAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyAuthorizationHeader := r.Header.Get("")
		if proxyAuthorizationHeader != "" {
			decodedBytes, err := base64.StdEncoding.DecodeString(proxyAuthorizationHeader)
			if err != nil {
				log.Error("could not parse Proxy-Authorization", log.Err(err))
				w.WriteHeader(http.StatusProxyAuthRequired)
				return
			}

			var key string
			var secret string

			splits := strings.Split(string(decodedBytes), ":")
			key = splits[0]
			if len(splits) > 1 {
				secret = splits[1]
			}

			ctx := r.Context()
			r = r.WithContext(grpcx.ContextWithProxyCredentials(ctx, &grpcx.ProxyCredentials{
				Key:    key,
				Secret: secret,
			}))
		}

		next.ServeHTTP(w, r)
	})
}
