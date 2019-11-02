package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type Deposit struct {
	id          int64  `json:"id"`
	Name        string `json:"name"`
	StatusColor string `json:"color"`
}

type DepositStatusesService struct {
	c *Client
}

// Do send request
func (s *DepositStatusesService) Do(ctx context.Context, opts ...RequestOption) ([]Deposit, error) {
	r := &request{
		method:   "GET",
		endpoint: "/public/deposit-statuses",
		secType:  secTypeNone,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []Deposit `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

type DepositStatusByIdService struct {
	c *Client

	id *int64
}

// Do send request
func (s *DepositStatusByIdService) Do(ctx context.Context, opts ...RequestOption) (*Deposit, error) {
	if s.id == nil {
		return nil, fmt.Errorf("id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/deposit-statuses/%d", *s.id),
		secType:  secTypeNone,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data Deposit `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *DepositStatusByIdService) Id(id int64) *DepositStatusByIdService {
	s.id = &id
	return s
}
