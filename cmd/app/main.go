package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	v1 "github.com/maxwell-ihiaso/theservicehood_core_service/api/v1"
	"github.com/maxwell-ihiaso/theservicehood_core_service/api/v1/authservice"
	"github.com/maxwell-ihiaso/theservicehood_core_service/config"
	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak"
	keycloakmiddleware "github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak/middleware"
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

	keycloakService := authservice.KeycloakService{Keycloak: kc}

	// Unprotected paths
	noAuthRouter.HandleFunc("/healthcheck", healthHandler).Methods(http.MethodGet)
	v1.InitAuth(router)
	// Protected Paths
	authRouter.HandleFunc("/getdcs", keycloakService.Getdocs).Methods(http.MethodGet)

	// apply middleware
	mdw := keycloakmiddleware.NewMiddleware(kc)
	authRouter.Use(mdw.VerifyToken)

	// create a server object
	s := &httpServer{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			Handler:      router,
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		},
	}

	fmt.Printf("\n====================\n\n")
	fmt.Printf("server created with address on %v:%v\n", host, port)
	fmt.Printf("\n====================\n")

	return s
}

func (s *httpServer) listen() error {
	fmt.Printf("\nserver listening on %v\n", s.server.Addr)
	return s.server.ListenAndServe()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\n==============================\n"))
	w.Write([]byte("running...\n"))
	w.Write([]byte("==============================\n"))
}

func main() {
	cfg := config.Load()

	s := newServer(cfg.Host, cfg.Port, keycloak.NewKeycloak())
	go s.listen()

	waitForShutdown(s.server)

}

func waitForShutdown(httpServer *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down HTTP server gracefully...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("\ncould not gracefully shutdown the server: %v\n", err)
	}

	fmt.Printf("\n====================\n\n")
	fmt.Printf("Stopped HTTP Server!\n")
	fmt.Printf("\n====================\n")

}
