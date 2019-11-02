package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type Withdrawal struct {
	id          int64  `json:"id"`
	Name        string `json:"name"`
	StatusColor string `json:"color"`
}

type WithdrawalStatusesService struct {
	c *Client
}

// Do send request
func (s *WithdrawalStatusesService) Do(ctx context.Context, opts ...RequestOption) ([]Withdrawal, error) {
	r := &request{
		method:   "GET",
		endpoint: "/public/withdrawal-statuses",
		secType:  secTypeNone,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []Withdrawal `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

type WithdrawalStatusByIdService struct {
	c *Client

	id *int64
}

// Do send request
func (s *WithdrawalStatusByIdService) Do(ctx context.Context, opts ...RequestOption) (*Withdrawal, error) {
	if s.id == nil {
		return nil, fmt.Errorf("id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/withdrawal-statuses/%d", *s.id),
		secType:  secTypeNone,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data Withdrawal `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *WithdrawalStatusByIdService) Id(id int64) *WithdrawalStatusByIdService {
	s.id = &id
	return s
}
