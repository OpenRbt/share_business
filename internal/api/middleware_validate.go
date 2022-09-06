// Code generated by mtgroup-generator.
package api

import (
	"net"
	"net/http"

	extauthapi "github.com/mtgroupit/mt-mock-extauthapi"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

const (
	headerValidate = extauthapi.ValidateHeaderName
)

// makeValidateToken creates middleware for checks validation token
func (svc *service) makeValidateToken(needValidate func(*http.Request) bool) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !needValidate(r) {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get(headerValidate)
			remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
			err := svc.checkerValidate(string(token), remoteIP)
			if err != nil {
				middlewareError(w, r, http.StatusForbidden, err.Error())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
