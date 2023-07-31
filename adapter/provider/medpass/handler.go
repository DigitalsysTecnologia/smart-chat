package medpass

import (
	"github.com/go-resty/resty/v2"
	"net/url"
	"smart-chat/internal/config"
)

type authorizer struct {
	BaseURL *url.URL
	Client  *resty.Client
}

type AuthorizerGateway interface {
	ValidateToken(token string) error
}

func NewAuthorizerGateway() AuthorizerGateway {

	cfg := config.NewConfigService()

	return &authorizer{
		BaseURL: &url.URL{
			Scheme: "https",
			Host:   cfg.Config.AuthorizerApiEndpoint,
			Path:   "/authorizer-v2/v2/",
		},
		Client: resty.New(),
	}
}
