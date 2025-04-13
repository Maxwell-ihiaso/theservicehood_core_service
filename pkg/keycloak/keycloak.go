package keycloak

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
)

type Keycloak struct {
	Gocloak      *gocloak.GoCloak //keycloak client
	ClientId     string           // clientId specified in keycloak
	ClientSecret string           // client secret specified in keycloak
	Realm        string           // realm specified in keycloak
}

// newKeycloak initializes and returns a pointer to a keycloak struct,
// setting up a GoCloak client and configuring it with environment
// variables for the Keycloak URL, client ID, client secret, and realm.
func NewKeycloak() *Keycloak {

	return &Keycloak{
		Gocloak:      gocloak.NewClient(os.Getenv("KEYCLOAK_URL")),
		ClientId:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		Realm:        os.Getenv("KEYCLOAK_REALM"),
	}
}
