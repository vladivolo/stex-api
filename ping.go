package stex

import (
	"context"
	"fmt"
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
	data, err := s.c.callAPI(ctx, r, opts...)

	fmt.Printf("PING: <%#s> <%s>\n", string(data), err)

	return err
}
