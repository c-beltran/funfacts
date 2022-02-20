package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rotisserie/eris"

	"github.com/c-beltran/funfacts/internal/facts"
)

type (
	// Client defines the APIs client.
	Client struct {
		host       string
		httpClient *http.Client
	}
)

func NewClient(client *http.Client, host string) *Client {
	return &Client{
		host:       host,
		httpClient: client,
	}
}

func (c *Client) FindDogFact(ctx context.Context) (facts.Diversity, error) {
	const path = `/api/facts`

	var response struct {
		Facts   []string `json:"facts"`
		Success bool     `json:"success"`
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, path), nil)
	if err != nil {
		return facts.Diversity{}, eris.Wrap(err, "creating request for dog facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Diversity{}, eris.Wrap(err, "request for dog facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Diversity{}, eris.Wrap(fmt.Errorf("bad status code from server %d", res.StatusCode), "")
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Diversity{}, eris.Wrap(err, "failed to parse response")
	}

	if !response.Success {
		return facts.Diversity{}, eris.Wrap(err, "failed to fetch, unsucessful")
	}

	return facts.Diversity{
		Dog: response.Facts[0],
	}, nil
}
