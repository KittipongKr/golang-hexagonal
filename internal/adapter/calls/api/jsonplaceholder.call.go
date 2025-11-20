package call_api

import (
	p "csat-servay/internal/core/port"

	"github.com/go-resty/resty/v2"
)

type jsonplaceholder struct {
	client resty.Client
}

func NewJsonplaceholderApi(client resty.Client) p.JsonplaceholderApi {
	return jsonplaceholder{client: client}
}

func (call jsonplaceholder) TestGetEndpoint() error {
	resp, err := call.client.R().
		Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return err
	}

	_ = resp

	return nil
}
