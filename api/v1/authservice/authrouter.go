package authservice

import (
	"net/http"

	"github.com/gorilla/mux"
	keycloakmiddleware "github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak/middleware"
)

func (c CoreAPI) mw(ct http.HandlerFunc) http.Handler {
	md := keycloakmiddleware.NewMiddleware(c.Keycloak)
	return md.VerifyToken(ct)
}

func (c CoreAPI) Router(router *mux.Router) {
	authRoute := router.PathPrefix("/v1/auth").Subrouter()

	authRoute.HandleFunc("/login-with-otp", c.HandleLoginTOTP).Methods("POST")
	authRoute.HandleFunc("/login", c.HandleLogin).Methods("POST")

	authRoute.Handle("/", c.mw(c.HandleLogin)).Methods("GET")

}
