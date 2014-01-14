package github

import (
	"github.com/42races/soauth/auth"
)

const (
	AuthUrl  = "https://github.com/login/oauth/authorize"
	TokenUrl = "https://github.com/login/oauth/access_token"
)

func Init(client_id, client_secret, redirect_url string, options map[string]string) *auth.Auth {
	auth := auth.Init(client_id, client_secret, redirect_url, AuthUrl, TokenUrl)
	return auth
}
