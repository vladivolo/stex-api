package stex

import (
	"context"
	"encoding/json"
)

type PairsGroup struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type PairsGroupsService struct {
	c *Client
}

// Do send request
func (s *PairsGroupsService) Do(ctx context.Context, opts ...RequestOption) ([]PairsGroup, error) {
	r := &request{
		method:   "GET",
		endpoint: "/public/pairs-goups",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []PairsGroup `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}
