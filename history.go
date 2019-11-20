package stex

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Trade struct {
	Id          int64     `json:"id"`
	BuyOrderId  int64     `json:"buy_order_id"`
	SellOrderId int64     `json:"sell_order_id"`
	Price       string    `json:"price"`
	Amount      string    `json:"amount"`
	TradeType   TradeType `json:"trade_type"`
	Timestamp   string    `json:"timestamp"`
}

type Fee struct {
	Id         int64   `json:"id"`
	CurrencyId int     `json:"currency_id"`
	Amount     float64 `json:"amount"`
	Timestamp  string  `json:"timestamp"`
}

type TradeOrderDetail struct {
	Id             int64   `json:""`
	CurrencyPairId int     `json:"currency_pair_id"`
	Price          string  `json:"price"`
	InitialAmount  string  `json:"initial_amount"`
	Type           string  `json:"type"`
	Created        string  `json:"created"`
	Timestamp      int64   `json:"timestamp"`
	Status         string  `json:"status"`
	Trades         []Trade `json:"trades"`
	Fees           []Fee   `json:"fees"`
}

type OrdersHistoryService struct {
	c *Client

	pair_id      *int
	order_status *OrderStatus
	time_start   *time.Time
	time_end     *time.Time
	limit        *int
	offset       *int
}

// Do send request
func (s *OrdersHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]OrderInfo, error) {
	r := &request{
		method:   "GET",
		endpoint: "/reports/orders",
		secType:  secTypeAPIKey,
	}

	if s.pair_id != nil {
		r.setParam("currencyPairId", *s.pair_id)
	}

	if s.time_start != nil {
		r.setParam("timeStart", s.time_start.Unix())
	}

	if s.time_end != nil {
		r.setParam("timeEnd", s.time_end.Unix())
	}

	if s.order_status != nil {
		r.setParam("orderStatus", *s.order_status)
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
		Data []OrderInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *OrdersHistoryService) CurrencyPairId(pair_id int) *OrdersHistoryService {
	s.pair_id = &pair_id
	return s
}

func (s *OrdersHistoryService) Status(status OrderStatus) *OrdersHistoryService {
	s.order_status = &status
	return s
}

func (s *OrdersHistoryService) TmStart(from time.Time) *OrdersHistoryService {
	s.time_start = &from
	return s
}

func (s *OrdersHistoryService) TmEnd(end time.Time) *OrdersHistoryService {
	s.time_end = &end
	return s
}

func (s *OrdersHistoryService) Limit(limit int) *OrdersHistoryService {
	s.limit = &limit
	return s
}

func (s *OrdersHistoryService) Offset(offset int) *OrdersHistoryService {
	s.offset = &offset
	return s
}

type TradesOrderHistoryService struct {
	c *Client

	order_id *int64
}

// Do send request
func (s *TradesOrderHistoryService) Do(ctx context.Context, opts ...RequestOption) (*TradeOrderDetail, error) {
	if s.order_id == nil {
		return nil, fmt.Errorf("order_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/reports/orders/%d", *s.order_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data TradeOrderDetail `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *TradesOrderHistoryService) OrderId(id int64) *TradesOrderHistoryService {
	s.order_id = &id
	return s
}

type CurrencyPairTradesHistoryService struct {
	c *Client

	pair_id    *int
	time_start *time.Time
	time_end   *time.Time
	limit      *int
	offset     *int
}

func (s *CurrencyPairTradesHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]Trade, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/reports/trades/%d", *s.pair_id),
		secType:  secTypeAPIKey,
	}

	if s.time_start != nil {
		r.setParam("timeStart", s.time_start.UTC().Unix())
	}

	if s.time_end != nil {
		r.setParam("timeEnd", s.time_end.UTC().Unix())
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
		Data []Trade `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *CurrencyPairTradesHistoryService) CurrencyPairId(pair_id int) *CurrencyPairTradesHistoryService {
	s.pair_id = &pair_id
	return s
}

func (s *CurrencyPairTradesHistoryService) TmStart(from time.Time) *CurrencyPairTradesHistoryService {
	s.time_start = &from
	return s
}

func (s *CurrencyPairTradesHistoryService) TmEnd(end time.Time) *CurrencyPairTradesHistoryService {
	s.time_end = &end
	return s
}

func (s *CurrencyPairTradesHistoryService) Limit(limit int) *CurrencyPairTradesHistoryService {
	s.limit = &limit
	return s
}

func (s *CurrencyPairTradesHistoryService) Offset(offset int) *CurrencyPairTradesHistoryService {
	s.offset = &offset
	return s
}
