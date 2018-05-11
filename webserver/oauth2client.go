package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	//NOTE(nick): reference article
	// - https://skarlso.github.io/2016/06/12/google-signin-with-go/
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Authorizer interface {
	LoadConfigurationFile(filePath string) error
	ConfigureGoogleAuth() error
}

type OAuth2Client struct {
	ID            string `json:"id"`
	Secret        string `json:"secret"`
	Configuration *oauth2.Config
}

func (client *OAuth2Client) LoadConfigurationFile(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return err
	}
	json.Unmarshal(file, &client)
	fmt.Printf("OAuth2 Client Loaded: %v\n", client)
	return nil
}

func (client *OAuth2Client) ConfigureGoogleAuth() error {
	*client = OAuth2Client{
		ID:     client.ID,
		Secret: client.Secret,
		Configuration: &oauth2.Config{
			ClientID:     client.ID,
			ClientSecret: client.Secret,
			RedirectURL:  "http://localhost:8080/auth",
			Scopes: []string{
				"",
			},
			Endpoint: google.Endpoint,
		},
	}
	return nil
}

var (
	GlobalOAuth2Client OAuth2Client
)

func init() {
	err := GlobalOAuth2Client.LoadConfigurationFile("./oauth2.json")
	if err != nil {
		os.Exit(1)
	}
}
