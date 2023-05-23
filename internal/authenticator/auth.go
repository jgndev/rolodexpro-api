package authenticator

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/jgndev/rolodexpro-api/internal/config"
	"golang.org/x/oauth2"
	"os"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator
func New() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+os.Getenv(config.Auth0Domain)+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv(config.Auth0ClientId),
		ClientSecret: os.Getenv(config.Auth0ClientSecret),
		RedirectURL:  os.Getenv(config.Auth0CCallbackUrl),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidConfig).Verify(ctx, rawIDToken)
}
