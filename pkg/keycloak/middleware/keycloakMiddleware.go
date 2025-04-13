package keycloakmiddleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak"
)

type KeyCloakMiddleware struct {
	keycloak *keycloak.Keycloak
}

func NewMiddleware(keycloak *keycloak.Keycloak) *KeyCloakMiddleware {
	return &KeyCloakMiddleware{keycloak: keycloak}
}

func (auth *KeyCloakMiddleware) extractToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func (auth *KeyCloakMiddleware) VerifyToken(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		// try to extract Authorization parameter from the HTTP header
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract the token
		token := auth.extractToken(bearerToken)

		if token == "" {
			http.Error(w, "Token missing", http.StatusUnauthorized)
			return
		}

		//// call Keycloak API to verify the access token
		result, err := auth.keycloak.Gocloak.RetrospectToken(context.Background(), token, auth.keycloak.ClientId, auth.keycloak.ClientSecret, auth.keycloak.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		jwt, _, err := auth.keycloak.Gocloak.DecodeAccessToken(context.Background(), token, auth.keycloak.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		jwtj, _ := json.Marshal(jwt)
		fmt.Printf("token: %v\n", string(jwtj))
		fmt.Printf("jwt: %+v\n\n", *jwt)
		fmt.Printf("result: %+v\n\n", *result)

		// check if the token isn't expired and valid
		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
