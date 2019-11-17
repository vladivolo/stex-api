package stex

import (
	"context"
	"encoding/json"
)

type Ping struct {
	Timestamp int64 `json:"server_timestamp"`
}

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (*Ping, error) {
	r := &request{
		method:   "GET",
		endpoint: "/public/ping",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)

	res := struct {
		APIError
		Data Ping `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}
