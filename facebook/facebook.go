package facebook

import (
	"encoding/json"
	"github.com/42races/soauth/auth"
	"net/http"
	"net/url"
)

type Graph struct {
	AccessToken string
}

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	Data []Item `json:"data"`
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

const (
	GraphHost = "graph.facebook.com"
	AuthUrl   = "https://graph.facebook.com/oauth/authorize"
	TokenUrl  = "https://graph.facebook.com/oauth/access_token"
)

func Init(client_id, client_secret, redirect_url string, options map[string]string) *auth.Auth {
	fba := auth.Init(client_id, client_secret, redirect_url, AuthUrl, TokenUrl)
	return fba
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

func (g *Graph) GetConnections(path string) ([]Item, error) {
	var data Data
	resp, err := http.Get(g.getRequestUri(path))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)

	return data.Data, nil
}
