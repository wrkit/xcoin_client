package xcoin_client

import (
	"fmt"
	"net/http"
	"net/url"
)

// client implements the Client interface
type client struct {
	cfg        *Config
	httpClient *http.Client
}

// NewClient creates a new XCoin bot client with the given configuration
func NewClient(config *Config) IClient {
	return &client{
		cfg:        config,
		httpClient: config.HTTPClient,
	}
}

func (c *client) requestUrl(path string) string {

	u, err := url.Parse(c.cfg.BaseURL)
	if err != nil {
		fmt.Printf("make api url %v\n", err)
	}

	u = u.JoinPath(path)
	return u.String()
}
