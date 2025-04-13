package authservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak"
	keycloakmodel "github.com/maxwell-ihiaso/theservicehood_core_service/pkg/keycloak/model"
)

type CoreAPI struct {
	router   *mux.Router
	Keycloak *keycloak.Keycloak
	usecase  ICore
}

func NewCoreAPI(router *mux.Router, usecase ICore, keycloak *keycloak.Keycloak) *CoreAPI {
	return &CoreAPI{router: router, usecase: usecase, Keycloak: keycloak}
}

func (c *CoreAPI) HandleLoginTOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &keycloakmodel.LoginResponse{}
	defer json.NewEncoder(w).Encode(res)

	rq := keycloakmodel.LoginTOTPRequest{}
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	err := rq.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	jwt, err := c.usecase.LoginTOTP(context.TODO(), c.Keycloak.ClientId, c.Keycloak.ClientSecret, c.Keycloak.Realm, rq.Username, rq.Password, rq.TOTP)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		res.Error = err.Error()
		return
	}

	res.Data.AccessToken = jwt.AccessToken
	res.Data.RefreshToken = jwt.RefreshToken
	res.Data.ExpiresIn = jwt.ExpiresIn

	w.WriteHeader(http.StatusOK)

}

func (c *CoreAPI) HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &keycloakmodel.LoginResponse{}
	defer json.NewEncoder(w).Encode(res)

	rq := keycloakmodel.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	err := rq.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	jwt, err := c.usecase.Login(context.TODO(), c.Keycloak.ClientId, c.Keycloak.ClientSecret, c.Keycloak.Realm, rq.Username, rq.Password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		res.Error = err.Error()
		return
	}

	res.Data.AccessToken = jwt.AccessToken
	res.Data.RefreshToken = jwt.RefreshToken
	res.Data.ExpiresIn = jwt.ExpiresIn

	w.WriteHeader(http.StatusOK)

}
