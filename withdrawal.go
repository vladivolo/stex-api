package stex

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Withdrawal struct {
	id          int64  `json:"id"`
	Name        string `json:"name"`
	StatusColor string `json:"color"`
}

type WithdrawalAdv struct {
	Id                 int64   `json:"id"`
	CurrencyId         int     `json:"currency_id"`
	CurrencyCode       string  `json:"currency_code"`
	Amount             string  `json:"amount"`
	Fee                string  `json:"fee"`
	FeeCurrencyId      int     `json:"fee_currency_id"`
	FeeCurrencyCode    string  `json:"fee_currency_code"`
	WithdrawalStatusId int     `json:"withdrawal_status_id"`
	Status             string  `json:"status"`
	StatusColor        string  `json:"status_color"`
	CreatedAt          string  `json:"created_at"`
	CreatedTs          string  `json:"created_ts"`
	UpdatedAt          string  `json:"updated_at"`
	UpdatedTs          string  `json:"updated_ts"`
	Txid               *string `json:"txid"`
	withdrawal_address Address `json:"withdrawal_address"`
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

type ProfileWithdrawalListService struct {
	c *Client

	currency_id *int64
	sort        *SortOrder
	time_start  *time.Time
	time_end    *time.Time
	limit       *int
	offset      *int
}

func (s *ProfileWithdrawalListService) Do(ctx context.Context, opts ...RequestOption) ([]WithdrawalAdv, error) {
	r := &request{
		method:   "GET",
		endpoint: "/profile/withdrawals",
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
		Data []WithdrawalAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *ProfileWithdrawalListService) CurrencyId(id int64) *ProfileWithdrawalListService {
	s.currency_id = &id
	return s
}

func (s *ProfileWithdrawalListService) Order(order SortOrder) *ProfileWithdrawalListService {
	s.sort = &order
	return s
}

func (s *ProfileWithdrawalListService) TmStart(from time.Time) *ProfileWithdrawalListService {
	s.time_start = &from
	return s
}

func (s *ProfileWithdrawalListService) TmEnd(end time.Time) *ProfileWithdrawalListService {
	s.time_end = &end
	return s
}

func (s *ProfileWithdrawalListService) Limit(limit int) *ProfileWithdrawalListService {
	s.limit = &limit
	return s
}

func (s *ProfileWithdrawalListService) Offset(offset int) *ProfileWithdrawalListService {
	s.offset = &offset
	return s
}

type ProfileWithdrawalInfoService struct {
	c *Client

	withdrawal_id *int64
}

func (s *ProfileWithdrawalInfoService) Do(ctx context.Context, opts ...RequestOption) (*WithdrawalAdv, error) {
	if s.withdrawal_id == nil {
		return nil, fmt.Errorf("withdrawal_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/profile/withdrawals/%d", *s.withdrawal_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data WithdrawalAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileWithdrawalInfoService) WithdrawalId(id int64) *ProfileWithdrawalInfoService {
	s.withdrawal_id = &id
	return s
}

type ProfileWithdrawalCreateService struct {
	c *Client

	currency_id                  *int64
	amount                       *float64
	address                      *string
	protocol_id                  *int
	additional_address_parameter *string // If withdrawal address requires the payment ID or some key or destination tag etc pass it here
}

func (s *ProfileWithdrawalCreateService) Do(ctx context.Context, opts ...RequestOption) (*WithdrawalAdv, error) {
	r := &request{
		method:   "POST",
		endpoint: "/profile/withdraw",
		secType:  secTypeAPIKey,
	}

	if s.currency_id != nil {
		r.setParam("currency_id", *s.currency_id)
	} else {
		return nil, fmt.Errorf("currency_id not init")
	}

	if s.amount != nil {
		r.setParam("amount", *s.amount)
	} else {
		return nil, fmt.Errorf("amount not init")
	}

	if s.address != nil {
		r.setParam("address", *s.address)
	} else {
		return nil, fmt.Errorf("withdraw address not init")
	}

	if s.protocol_id != nil {
		r.setParam("protocol_id", s.protocol_id)
	}

	if s.additional_address_parameter != nil {
		r.setParam("additional_address_parameter", *s.additional_address_parameter)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data WithdrawalAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileWithdrawalCreateService) CurrencyId(id int64) *ProfileWithdrawalCreateService {
	s.currency_id = &id
	return s
}

func (s *ProfileWithdrawalCreateService) Amount(amount float64) *ProfileWithdrawalCreateService {
	s.amount = &amount
	return s
}

func (s *ProfileWithdrawalCreateService) Address(address string) *ProfileWithdrawalCreateService {
	s.address = &address
	return s
}

func (s *ProfileWithdrawalCreateService) ProtocolId(id int) *ProfileWithdrawalCreateService {
	s.protocol_id = &id
	return s
}

func (s *ProfileWithdrawalCreateService) PaymentId(id string) *ProfileWithdrawalCreateService {
	s.additional_address_parameter = &id
	return s
}

type ProfileWithdrawalCancelService struct {
	c *Client

	withdrawal_id *int64
}

func (s *ProfileWithdrawalCancelService) Do(ctx context.Context, opts ...RequestOption) (*WithdrawalAdv, error) {
	if s.withdrawal_id == nil {
		return nil, fmt.Errorf("withdrawal_id not init")
	}

	r := &request{
		method:   "DELETE",
		endpoint: fmt.Sprintf("/profile/withdraw/%d", *s.withdrawal_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data WithdrawalAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileWithdrawalCancelService) WithdrawalId(id int64) *ProfileWithdrawalCancelService {
	s.withdrawal_id = &id
	return s
}
