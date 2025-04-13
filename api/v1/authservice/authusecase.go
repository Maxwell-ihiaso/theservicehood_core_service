package authservice

import (
	"context"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak"
)

type KeycloakService struct {
	Keycloak *keycloak.Keycloak
}

func NewKeycloakService(keycloak *keycloak.Keycloak) ICore {
	return &KeycloakService{Keycloak: keycloak}
}

func (ks *KeycloakService) Login(ctx context.Context, clientID string, clientSecret, realm, username, password string) (*gocloak.JWT, error) {
	return ks.Keycloak.Gocloak.Login(ctx, clientID, clientSecret, realm, username, password)
}

func (ks *KeycloakService) LoginTOTP(ctx context.Context, clientID string, clientSecret, realm, username, password, totp string) (*gocloak.JWT, error) {
	return ks.Keycloak.Gocloak.LoginOtp(context.TODO(), clientID, clientSecret, realm, username, password, totp)
}

func (ks *KeycloakService) Getdocs(w http.ResponseWriter, r *http.Request) {}
