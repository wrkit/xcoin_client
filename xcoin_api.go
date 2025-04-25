package xcoin_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// GetVersion implements the version endpoint with internal non-blocking execution
func (c *client) GetVersion(ctx context.Context) (*VersionResponse, error) {
	url := c.requestUrl("Ver")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestCreation, err)
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestExecution, err)
	}
	defer resp.Body.Close()

	var version *VersionResponse

	if resp.Header.Get("Content-Type") == "text/html" {
		version, err = parseVersionResponse(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrHTMLParsing, err)
		}
		version.BotAddress = c.cfg.getBotIpPost()
	} else {
		version = &VersionResponse{}
		if err = json.NewDecoder(resp.Body).Decode(version); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrJSONDecoding, err)
		}
		version.BotAddress = c.cfg.getBotIpPost()
	}

	return version, nil
}

func (c *client) GetPairList2(ctx context.Context) ([]Pair, error) {
	url := c.requestUrl("GetListPair2")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestCreation, err)
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestExecution, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != StatusOK {
		return nil, fmt.Errorf("%w: got %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	var pairs []Pair
	if err := json.NewDecoder(resp.Body).Decode(&pairs); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrJSONDecoding, err)
	}

	return pairs, nil
}

// GetPairCurrentSettings retrieves current settings for a trading pair
func (c *client) GetPairCurrentSettings(ctx context.Context, keyWork int) (*KeyValueResponse, error) {
	url := fmt.Sprintf("%s/%d", c.requestUrl("PairCurSetting"), keyWork)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var settings KeyValueResponse
	if err := json.NewDecoder(resp.Body).Decode(&settings); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &settings, nil
}

// GetPairOrders retrieves current orders for a trading pair
func (c *client) GetPairOrders(ctx context.Context, keyWork int) ([]Order, error) {
	url := fmt.Sprintf("%s/%d", c.requestUrl("PairCurOrders"), keyWork)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var orders []Order
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return orders, nil
}

// AddPair adds a new trading pair
func (c *client) AddPair(ctx context.Context, val1, val2, birga, mode string) error {
	url := fmt.Sprintf("%s/%s/%s/%s/%s", c.requestUrl("PairAdd"), val1, val2, birga, mode)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrRequestCreation, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrRequestExecution, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != StatusOK {
		return fmt.Errorf("%w: got %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	return nil
}

// UpdatePairSettings updates settings for a trading pair
func (c *client) UpdatePairSettings(ctx context.Context, keyWork int, settings *KeyValueResponse) error {
	url := fmt.Sprintf("%s/%d", c.requestUrl("PairCurSetting"), keyWork)

	body, err := json.Marshal(settings)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *client) ExecuteCommand(ctx context.Context, keyWork int, command CommandType) error {
	url := fmt.Sprintf("%s/%d/%s", c.requestUrl("ExecuteCommand"), keyWork, command)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrRequestCreation, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrRequestExecution, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != StatusOK {
		return fmt.Errorf("%w: got %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	return nil
}

// parseVersionResponse parses the HTML response from the version endpoint
func parseVersionResponse(body io.Reader) (*VersionResponse, error) {

	doc, err := html.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	response := &VersionResponse{}
	response.Pairs = make([]PairInfo, 0)

	// Helper function to extract text content from a node
	getText := func(n *html.Node) string {
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			return strings.TrimSpace(n.FirstChild.Data)
		}
		return ""
	}

	// Helper function to traverse the HTML tree
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "tr":
				// Parse table rows for pair information
				var keyWork, pair string
				for td := n.FirstChild; td != nil; td = td.NextSibling {
					if td.Type == html.ElementNode && td.Data == "td" {
						text := getText(td)
						if keyWork == "" {
							keyWork = text
						} else {
							pair = text
							if keyWork != "keyWork" { // Skip header row
								kw, _ := strconv.Atoi(keyWork)
								response.Pairs = append(response.Pairs, PairInfo{KeyWork: kw, Pair: pair})
							}
							break
						}
					}
				}
			case "span":
				// Extract bot ID and address
				for _, attr := range n.Attr {
					if attr.Key == "style" && attr.Val == "color: blue" {
						text := getText(n)
						if strings.HasPrefix(text, "http") {
							response.BotAddress = text
						} else {
							response.BotID = text
						}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	response.PairCount = len(response.Pairs)

	return response, nil
}

// Helper function to execute HTTP requests with non-blocking behavior
func (c *client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	respCh := make(chan struct {
		resp *http.Response
		err  error
	}, 1)

	go func() {
		resp, err := c.httpClient.Do(req)
		respCh <- struct {
			resp *http.Response
			err  error
		}{resp, err}
	}()

	select {
	case result := <-respCh:
		return result.resp, result.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
