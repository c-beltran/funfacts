package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rotisserie/eris"

	"github.com/c-beltran/funfacts/internal/facts"
)

func (c *Client) FindCatFact(ctx context.Context) (facts.Cat, error) {
	var response struct {
		Facts []string `json:"data"`
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, ""), nil)
	if err != nil {
		return facts.Cat{}, eris.Wrap(err, "creating request for cat facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Cat{}, eris.Wrap(err, "request for cat facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Cat{}, eris.Wrapf(err, "bad status code from server %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Cat{}, eris.Wrap(err, "failed to parse response")
	}

	return facts.Cat{
		Fact: response.Facts[0],
	}, nil
}
