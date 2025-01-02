package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/utils/keycloak"
)

type KeycloakService struct {
	Keycloak *keycloak.Keycloak
}

type LoginResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
	Error   string            `json:"error"`
}

type LoginResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type LoginRequest struct {
	Username string `json :"username"`
	Password string `json:"password"`
}

func (ks *KeycloakService) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &LoginResponse{}
	defer json.NewEncoder(w).Encode(res)

	rq := LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwt, err := ks.Keycloak.Gocloak.Login(context.Background(), ks.Keycloak.ClientId, ks.Keycloak.ClientSecret, ks.Keycloak.Realm, rq.Username, rq.Password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		res.Error = err.Error()
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	res.Data.AccessToken = jwt.AccessToken
	res.Data.RefreshToken = jwt.RefreshToken
	res.Data.ExpiresIn = jwt.ExpiresIn

	w.WriteHeader(http.StatusOK)

}

func (ks *KeycloakService) Getdocs(w http.ResponseWriter, r *http.Request) {}
