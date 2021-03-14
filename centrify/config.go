package centrify

import (
	"fmt"

	"github.com/marcozj/golang-sdk/dmc"
	"github.com/marcozj/golang-sdk/oauth"
)

// Config - Centrify Vault client struct
type Config struct {
	URL            string
	AppID          string
	Scope          string
	Username       string
	Password       string
	Token          string
	UseDMC         bool
	LogLevel       string
	LogPath        string
	SkipCertVerify bool
}

// Valid - Validate provider configuration
func (c *Config) Valid() error {
	if c.URL == "" {
		return fmt.Errorf("Tenant URL must be provided for the CentrifyVault provider")
	}
	if c.Scope == "" {
		return fmt.Errorf("Scope must be provided for the CentrifyVault provider")
	}

	if !c.UseDMC && c.Token == "" {
		// If DMC isn't used and token isn't supplied, make sure appid user username is provided
		if c.AppID == "" {
			return fmt.Errorf("AppID must be provided for the CentrifyVault provider")
		}
		if c.Username == "" {
			return fmt.Errorf("Username or token must be provided for the CentrifyVault provider")
		}
	}

	return nil
}

func (c *Config) getClient() (interface{}, error) {
	var client interface{}
	var err error
	if c.UseDMC {
		// use DMC to return authenticated Rest client
		call := dmc.DMC{}
		call.Service = c.URL
		call.Scope = c.Scope
		call.Token = c.Token
		call.SkipCertVerify = c.SkipCertVerify

		client, err = call.GetClient()
	} else {
		// use OAuth authentication
		call := oauth.OauthClient{
			Service:        c.URL,
			AppID:          c.AppID,
			Scope:          c.Scope,
			Token:          c.Token,
			ClientID:       c.Username,
			ClientSecret:   c.Password,
			SkipCertVerify: c.SkipCertVerify,
		}
		client, err = call.GetClient()
	}
	return client, err
}
