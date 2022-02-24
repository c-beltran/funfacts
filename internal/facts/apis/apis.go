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

// FindCatFact returns a cat fact.
func (c *Client) FindCatFact(ctx context.Context) (facts.Topic, error) {
	var response struct {
		Facts []string `json:"data"`
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, ""), nil)
	if err != nil {
		return facts.Topic{}, eris.Wrap(err, "creating request for cat facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "request for cat facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Topic{}, eris.Wrap(fmt.Errorf("bad status code from server %d", res.StatusCode), "")
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Topic{}, eris.Wrap(err, "failed to parse response")
	}

	return facts.Topic{
		Cat: response.Facts[0],
	}, nil
}

// FindDogFact returns a dog fact.
func (c *Client) FindDogFact(ctx context.Context) (facts.Topic, error) {
	const path = `/api/facts`

	var response struct {
		Facts   []string `json:"facts"`
		Success bool     `json:"success"`
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, path), nil)
	if err != nil {
		return facts.Topic{}, eris.Wrap(err, "creating request for dog facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "request for dog facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Topic{}, eris.Wrap(fmt.Errorf("bad status code from server %d", res.StatusCode), "")
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Topic{}, eris.Wrap(err, "failed to parse response")
	}

	if !response.Success {
		return facts.Topic{}, eris.Wrap(err, "failed to fetch, unsucessful")
	}

	return facts.Topic{
		Dog: response.Facts[0],
	}, nil
}

// FindEntertainmentFact returns a entertainment fact.
func (c *Client) FindEntertainmentFact(ctx context.Context) (facts.Topic, error) {
	type (
		data struct {
			ID       string `json:"id"`
			Fact     string `json:"fact"`
			Category string `json:"cat"`
		}
	)

	var response struct {
		Data data `json:"data"`
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, ""), nil)
	if err != nil {
		return facts.Topic{}, eris.Wrap(err, "creating request for entertainment facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "request for entertainment facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Topic{}, eris.Wrap(fmt.Errorf("bad status code from server %d", res.StatusCode), "")
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Topic{}, eris.Wrap(err, "failed to parse response")
	}

	return facts.Topic{
		Entertainment: response.Data.Fact,
	}, nil
}

// FindTrivialFact returns a trivial fact.
func (c *Client) FindTrivialFact(ctx context.Context) (facts.Topic, error) {
	var response struct {
		Fact string `json:"text"`
	}

	path := "/random.json?language=en"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, path), nil)
	if err != nil {
		return facts.Topic{}, eris.Wrap(err, "creating request for trivial facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "request for trivial facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Topic{}, eris.Wrap(fmt.Errorf("bad status code from server %d", res.StatusCode), "")
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Topic{}, eris.Wrap(err, "failed to parse response")
	}

	return facts.Topic{
		Trivial: response.Fact,
	}, nil
}
