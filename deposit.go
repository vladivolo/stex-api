package stex

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Deposit struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	StatusColor string `json:"color"`
}

type DepositAdv struct {
	Id                     int64   `json:"id"`
	CurrencyId             int     `json:"currency_id"`
	CurrencyCode           string  `json:"currency_code"`
	DepositFeeCurrencyId   int     `json:"deposit_fee_currency_id"`
	DepositFeeCurrencyCode string  `json:"deposit_fee_currency_code"`
	Amount                 float64 `json:"amount"`
	Fee                    float64 `json:"fee"`
	Txid                   string  `json:"txid"`
	DepositStatusId        int     `json:"deposit_status_id"`
	Status                 string  `json:"status"`
	StatusColor            string  `json:"status_color"`
	CreatedAt              string  `json:"created_at"`
	Timestamp              int64   `json:"timestamp"`
	Confirmations          string  `json:"confirmations"`
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

type ProfileDepositsListService struct {
	c *Client

	currency_id *int64
	sort        *SortOrder
	time_start  *time.Time
	time_end    *time.Time
	limit       *int
	offset      *int
}

func (s *ProfileDepositsListService) Do(ctx context.Context, opts ...RequestOption) ([]DepositAdv, error) {
	r := &request{
		method:   "GET",
		endpoint: "/profile/deposits",
		secType:  secTypeAPIKey,
	}

	if s.currency_id != nil {
		r.setParam("currencyId", *s.currency_id)
	}

	if s.sort != nil {
		r.setParam("sort", *s.sort)
	}

	if s.time_start != nil {
		r.setParam("timeStart", s.time_start.Unix())
	}

	if s.time_end != nil {
		r.setParam("timeEnd", s.time_end.Unix())
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
		Data []DepositAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *ProfileDepositsListService) CurrencyId(id int64) *ProfileDepositsListService {
	s.currency_id = &id
	return s
}

func (s *ProfileDepositsListService) Order(order SortOrder) *ProfileDepositsListService {
	s.sort = &order
	return s
}

func (s *ProfileDepositsListService) TmStart(from time.Time) *ProfileDepositsListService {
	s.time_start = &from
	return s
}

func (s *ProfileDepositsListService) TmEnd(end time.Time) *ProfileDepositsListService {
	s.time_end = &end
	return s
}

func (s *ProfileDepositsListService) Limit(limit int) *ProfileDepositsListService {
	s.limit = &limit
	return s
}

func (s *ProfileDepositsListService) Offset(offset int) *ProfileDepositsListService {
	s.offset = &offset
	return s
}

type ProfileDepositInfoService struct {
	c *Client

	deposit_id *int64
}

func (s *ProfileDepositInfoService) Do(ctx context.Context, opts ...RequestOption) (*DepositAdv, error) {
	if s.deposit_id == nil {
		return nil, fmt.Errorf("deposit_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/profile/deposits/%d", *s.deposit_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data DepositAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileDepositInfoService) DepositId(id int64) *ProfileDepositInfoService {
	s.deposit_id = &id
	return s
}
