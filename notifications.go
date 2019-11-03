package stex

import (
	"context"
	"encoding/json"
)

type Notification struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Date  string `json:"date"`
}

type ProfileNotificationsService struct {
	c *Client

	limit  *int
	offset *int
}

// Do send request
func (s *ProfileNotificationsService) Do(ctx context.Context, opts ...RequestOption) ([]Notification, error) {
	r := &request{
		method:   "GET",
		endpoint: "/profile/notifications",
		secType:  secTypeAPIKey,
	}

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []Notification `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *ProfileNotificationsService) Limit(limit int) *ProfileNotificationsService {
	s.limit = &limit
	return s
}

func (s *ProfileNotificationsService) Offset(offset int) *ProfileNotificationsService {
	s.offset = &offset
	return s
}
