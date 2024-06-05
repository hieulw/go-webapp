package auth

import (
	"log"
	"ticket/internal/config"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/openidConnect"
)

type Auth struct{}

func New(cfg config.Config) {
	openidConnect, err := openidConnect.New(
		cfg.Auth.ClientID,
		cfg.Auth.ClientSecret,
		cfg.Auth.CallbackURL,
		cfg.Auth.DiscoveryEndpoint,
	)
	if err != nil {
		log.Fatal(err)
	}
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}
}
