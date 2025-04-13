package v1

import (
	"github.com/gorilla/mux"
	"github.com/maxwell-ihiaso/theservicehood_core_service/api/v1/authservice"
	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak"
)

func InitAuth(router *mux.Router) authservice.KeycloakService {
	kc := keycloak.NewKeycloak()
	keycloakService := authservice.KeycloakService{Keycloak: kc}
	api := authservice.NewCoreAPI(router, &keycloakService, kc)
	api.Router(router)

	return keycloakService
}
