package client

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var client *http.Client

// InitialiseClient initialises the OAuth2 client for authenticating with account api
func InitialiseClient(clientID, clientSecret, tokenURL string) {
	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"accounts.read", "accounts.create"},
	}

	config.Token(oauth2.NoContext)

	client = config.Client(oauth2.NoContext)
}
