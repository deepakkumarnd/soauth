package oauth_fb

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type FBAuth struct {
	oauth.Config
	Token *oauth.Token
}

type Graph struct {
	AccessToken string
}

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Profile struct {
	Id                  string  `json:"id"`
	Username            string  `json:"username"`
	Name                string  `json:"name"`
	FirstName           string  `json:"first_name"`
	LastName            string  `json:"last_name"`
	Link                string  `json:"link"`
	Gender              string  `json:"gender"`
	Timezone            float32 `json:"timezone"`
	Locale              string  `json:"locale"`
	Verified            bool    `json:"verified"`
	UpdatedTime         string  `json:"updated_time"`
	Hometown            Item    `json:"hometown"`
	Location            Item    `json:"location"`
	Sports              []Item  `json:"sports"`
	FavoriteTeams       []Item  `json:"favourite_teams"`
	Languages           []Item  `json:"languages"`
	InspirationalPeople []Item  `json:"inspirational_people"`
}

const GraphHost = "graph.facebook.com"

func Init(client_id, client_secret, redirect_url string, options map[string]string) *FBAuth {
	fba := new(FBAuth)
	fba.ClientId = client_id
	fba.ClientSecret = client_secret
	fba.RedirectURL = redirect_url
	fba.AuthURL = "https://graph.facebook.com/oauth/authorize"
	fba.TokenURL = "https://graph.facebook.com/oauth/access_token"

	return fba
}

func (fba *FBAuth) LoginURL() string {
	// demo is a random string to prevent cross site requests
	return fba.AuthCodeURL("demo")
}

func (fba *FBAuth) Authorize(code string) (string, error) {
	t := &oauth.Transport{Config: &fba.Config}
	tok, err := t.Exchange(code)
	fba.Token = tok
	if err != nil {
		fmt.Println("Error in getting token")
	}

	return fba.AccessToken(), err
}

func (fba *FBAuth) AccessToken() string {
	var token string
	if fba.Token != nil {
		token = fba.Token.AccessToken
	}
	return token
}

func (g *Graph) getRequestUri(path string) string {
	var uri url.URL
	query := url.Values{"access_token": {g.AccessToken}}.Encode()

	uri.Host = GraphHost
	uri.Path = path
	uri.Scheme = "https"
	uri.RawQuery = query

	return uri.String()
}

func (g *Graph) GetObject(object string) (*Profile, error) {
	var profile Profile
	resp, err := http.Get(g.getRequestUri(object))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&profile)
	return &profile, nil
}
