package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rotisserie/eris"

	"github.com/c-beltran/funfacts/internal/facts"
)

func (c *Client) FindTrivialFact(ctx context.Context) (facts.Trivial, error) {
	var response struct {
		Fact string `json:"text"`
	}

	path := "/random.json?language=en"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, path), nil)
	if err != nil {
		return facts.Trivial{}, eris.Wrap(err, "creating request for trivial facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Trivial{}, eris.Wrap(err, "request for trivial facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Trivial{}, eris.Wrapf(err, "bad status code from server %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Trivial{}, eris.Wrap(err, "failed to parse response")
	}

	return facts.Trivial{
		Fact: response.Fact,
	}, nil
}
