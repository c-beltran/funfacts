package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rotisserie/eris"

	"github.com/c-beltran/funfacts/internal/facts"
)

func (c *Client) FindEntertainmentFact(ctx context.Context) (facts.Entertainment, error) {
	type (
		data struct {
			ID                 string `json:"id"`
			Fact               string `json:"fact"`
			Entertainmentegory string `json:"cat"`
		}
	)

	var response struct {
		Data data `json:"data"`
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.host, ""), nil)
	if err != nil {
		return facts.Entertainment{}, eris.Wrap(err, "creating request for entertainment facts failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if err != nil {
			return facts.Entertainment{}, eris.Wrap(err, "request for entertainment facts failed")
		}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return facts.Entertainment{}, eris.Wrapf(err, "bad status code from server %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return facts.Entertainment{}, eris.Wrap(err, "failed to parse response")
	}

	return facts.Entertainment{
		Fact: response.Data.Fact,
	}, nil
}
