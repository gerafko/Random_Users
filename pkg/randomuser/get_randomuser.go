package randomuser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func (c *Client) RandomUserGET(ctx context.Context) (*GetRandomUserResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error new request witch context")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error http client do")
	}

	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusOK:
		users := &GetRandomUserResponse{}
		if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
			return nil, errors.Wrap(err, "error decode body")
		}

		return users, nil
	case http.StatusNoContent:
		return nil, ErrNoUsers
	default:
		errText := ""
		if res.Body != http.NoBody {
			resBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, errors.Wrap(err, "error read body")
			}
			errText = string(resBytes)
		}

		return nil, errors.Wrap(errors.New("error request"), fmt.Sprintf("error: service return status"+
			" code %v, body: %v", res.StatusCode, errText))
	}
}
