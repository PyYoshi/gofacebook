package gofacebook

import (
	"encoding/json"
	"github.com/robfig/goauth2/oauth"
)

// Facebook OAuth
const (
	FbAuthURL      = "https://graph.facebook.com/oauth/authorize"
	FbTokenURL     = "https://graph.facebook.com/oauth/access_token"
	FbRedirectURL  = "http://idemodev.com:9000/App/Auth"
	FbGraphAPIHost = "graph.facebook.com"
)

type Config struct {
	ClientId     string
	ClientSecret string
	RedirectURL  string
	AccessToken  string
	Scope        string
}

type OAuth struct {
	FacebookConfig Config
	FbOAuth        *oauth.Config
}

func (o *OAuth) NewOAuth(config Config) *OAuth {
	return &OAuth{
		FacebookConfig: config,
		FbOAuth: &oauth.Config{
			ClientId:     Config.ClientId,
			ClientSecret: Config.ClientSecret,
			AuthURL:      FbAuthURL,
			TokenURL:     FbTokenURL,
			RedirectURL:  FbRedirectURL,
			Scope:        Config.Scope,
		},
	}
}

func (o *OAuth) GetAuthURL() {
	return o.FbOAuth.AuthCodeURL("foo")
}

func (o *OAuth) GetAccessToken(code string) oauth.Token {
	t := &oauth.Transport{Config: o.FbOAuth}
	tok, err := t.Exchange(code)
	if err != nil {
		// TODO:
	}
	return tok
}

type GraphAPI struct {
	FacebookConfig Config
	FbOAuth        *oauth.Config
	SSL            bool
}

func (g *GraphAPI) _build_url() string {
	var url string
	// SSL
	scheme := "http://"
	if g.SSL != nil && g.SSL {
		scheme = "https://"
	}

	//
	url = scheme + FbGraphAPIHost

}

func (g *GraphAPI) _get() string {
	return ""
}

func (g *GraphAPI) _post() string {
	return ""
}

func (g *GraphAPI) GetMe() {

}
