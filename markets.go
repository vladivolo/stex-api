package stex

import (
	"context"
	"encoding/json"
)

type MarketInfo struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AvailableMarketsService struct {
	c *Client
}

// Do send request
func (s *AvailableMarketsService) Do(ctx context.Context, opts ...RequestOption) ([]MarketInfo, error) {
	r := &request{
		method:   "GET",
		endpoint: "/public/markets",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []MarketInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}
