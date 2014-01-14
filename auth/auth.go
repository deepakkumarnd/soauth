package auth

import (
	"code.google.com/p/goauth2/oauth"
)

type Auth struct {
	oauth.Config
	Token *oauth.Token
}

func Init(client_id, client_secret, redirect_url, auth_url, token_url string) *Auth {
	auth := new(Auth)
	auth.ClientId = client_id
	auth.ClientSecret = client_secret
	auth.RedirectURL = redirect_url
	auth.AuthURL = auth_url
	auth.TokenURL = token_url

	return auth
}

func (auth *Auth) LoginURL() string {
	// TODO: generate a random string and validate it
	// demo is a random string to prevent cross site requests
	return auth.AuthCodeURL("demo")
}

func (auth *Auth) Authorize(code string) (string, error) {
	t := &oauth.Transport{Config: &auth.Config}
	tok, err := t.Exchange(code)
	auth.Token = tok
	return auth.AccessToken(), err
}

func (auth *Auth) AccessToken() string {
	var token string
	if auth.Token != nil {
		token = auth.Token.AccessToken
	}
	return token
}
