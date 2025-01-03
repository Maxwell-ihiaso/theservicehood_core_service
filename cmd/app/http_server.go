package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/maxwell-ihiaso/theservicehood_core_service/api/v1/service"
	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/utils/keycloak"
)

type httpServer struct {
	server *http.Server
}

func newServer(host, port string, kc *keycloak.Keycloak) *httpServer {

	// create a root router
	router := mux.NewRouter()

	// add a subrouter based on matcher func
	// note, routers are processed one by one in order, so that if one of the routing matches other won't be processed
	noAuthRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.Header.Get("Authorization") == ""
	}).Subrouter()

	// add one more subrouter for the authenticated service methods
	authRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()

	keycloakService := service.KeycloakService{Keycloak: kc}

	// Unprotected paths
	noAuthRouter.HandleFunc("/healthcheck", healthHandler).Methods(http.MethodGet)
	noAuthRouter.HandleFunc("/login", keycloakService.Login).Methods(http.MethodPost)

	// Protected Paths
	authRouter.HandleFunc("/getdcs", keycloakService.Getdocs).Methods(http.MethodGet)

	// apply middleware
	mdw := keycloak.NewMiddleware(kc)
	authRouter.Use(mdw.VerifyToken)

	// create a server object
	s := &httpServer{
		server: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", host, port),
			Handler:      router,
			WriteTimeout: time.Hour,
			ReadTimeout:  time.Hour,
		},
	}

	return s
}

func (s *httpServer) listen() error {
	return s.server.ListenAndServe()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\n==============================\n"))
	w.Write([]byte("running...\n"))
	w.Write([]byte("==============================\n"))
}
