package foursqure

import (
	"github.com/42races/soauth/auth"
)

const (
	AuthUrl  = "https://foursquare.com/oauth2/authorize"
	TokenUrl = "https://foursquare.com/oauth2/access_token"
)

func Init(client_id, client_secret, redirect_url string, options map[string]string) *auth.Auth {
	fsa := auth.Init(client_id, client_secret, redirect_url, AuthUrl, TokenUrl)
	return fsa
}
