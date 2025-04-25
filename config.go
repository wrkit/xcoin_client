package xcoin_client

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	BaseURL    string
	HTTPClient *http.Client
	Timeout    time.Duration
}

func NewConfig(botUrl string) *Config {

	timeout := 120 * time.Second
	httpClient := &http.Client{
		Timeout: timeout,
	}

	c := &Config{
		BaseURL:    "",
		HTTPClient: httpClient,
		Timeout:    timeout,
	}
	c.BaseURL = c.makeUrl(botUrl, "XCoin")
	return c
}

func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	c.HTTPClient.Timeout = timeout
	return c
}

func (c *Config) WithBotUrlPath(suffix string) *Config {
	c.BaseURL = c.makeUrl(c.BaseURL, suffix)
	return c
}

func (c *Config) makeUrl(botUrl, botUrlPath string) string {

	u, err := url.Parse(botUrl)
	if err != nil {
		log.Fatalf("Invalid bot url: %v", err)
	}
	if strings.HasSuffix(botUrl, botUrlPath) {
		return botUrl
	}

	u = u.JoinPath(botUrlPath)
	return u.String()
}

func (c *Config) getBotIpPost() string {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		log.Fatalf("Invalid bot url: %v", err)
	}

	return fmt.Sprintf("%s:%s", u.Hostname(), u.Port())
}
