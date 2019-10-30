package stex

import (
	"context"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   "GET",
		endpoint: "/public/ping",
		secType:  secTypeNone,
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	return err
}
