package keycloakmodel

import "fmt"

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

type LoginTOTPRequest struct {
	Username string `json :"username"`
	Password string `json:"password"`
	TOTP     string `json:"totp"`
}

func (lr *LoginRequest) Validate() error {
	if lr.Username == "" {
		return fmt.Errorf("username is required")
	}
	if lr.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (lrt *LoginTOTPRequest) Validate() error {
	if lrt.Username == "" {
		return fmt.Errorf("username is required")
	}
	if lrt.Password == "" {
		return fmt.Errorf("password is required")
	}
	if lrt.TOTP == "" {
		return fmt.Errorf("totp is required")
	}
	return nil
}
