package authservice

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type (
	ICore interface {
		Login(ctx context.Context, clientID string, clientSecret, realm, username, password string) (*gocloak.JWT, error)
		LoginTOTP(ctx context.Context, clientID string, clientSecret, realm, username, password, totp string) (*gocloak.JWT, error)
	}
)
