package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type OrderInfo struct {
	Id              int64     `json:"id"`
	CurrencyPairId  int       `json:"currency_pair_id"`
	Price           string    `json:"price"`
	TriggerPrice    float64   `json:"trigger_price"`
	InitialAmount   string    `json:"initial_amount"`
	ProcessedAmount string    `json:"processed_amount"`
	Type            OrderType `json:"type"`
	OriginalType    OrderType `json:"original_type"`
	Created         string    `json:"created"`
	Timestamp       int64     `json:"timestamp"`
	Status          string    `json:"status"`
}

type DeletedOrders struct {
	Processing []OrderInfo `json:"put_into_processing_queue"`
	Pending    []OrderInfo `json:"not_put_into_processing_queue"`
	Message    string      `json:"message"`
}

type OpenOrdersListService struct {
	c *Client

	limit  *int
	offset *int
}

// Do send request
func (s *OpenOrdersListService) Do(ctx context.Context, opts ...RequestOption) ([]OrderInfo, error) {
	r := &request{
		method:   "GET",
		endpoint: "/trading/orders",
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
		Data []OrderInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *OpenOrdersListService) Limit(limit int) *OpenOrdersListService {
	s.limit = &limit
	return s
}

func (s *OpenOrdersListService) Offset(offset int) *OpenOrdersListService {
	s.offset = &offset
	return s
}

type OpenOrdersDeleteService struct {
	c *Client
}

// Do send request
func (s *OpenOrdersDeleteService) Do(ctx context.Context, opts ...RequestOption) (*DeletedOrders, error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/trading/orders",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data DeletedOrders `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

type CurrencyPairOpenOrdersListService struct {
	c *Client

	pair_id *int
	limit   *int
	offset  *int
}

// Do send request
func (s *CurrencyPairOpenOrdersListService) Do(ctx context.Context, opts ...RequestOption) ([]OrderInfo, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not set")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/trading/orders/%d", *s.pair_id),
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
		Data []OrderInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *CurrencyPairOpenOrdersListService) CurrencyPairId(pair_id int) *CurrencyPairOpenOrdersListService {
	s.pair_id = &pair_id
	return s
}

func (s *CurrencyPairOpenOrdersListService) Limit(limit int) *CurrencyPairOpenOrdersListService {
	s.limit = &limit
	return s
}

func (s *CurrencyPairOpenOrdersListService) Offset(offset int) *CurrencyPairOpenOrdersListService {
	s.offset = &offset
	return s
}

type CurrencyPairOpenOrdersDeleteService struct {
	c *Client

	pair_id *int
}

// Do send request
func (s *CurrencyPairOpenOrdersDeleteService) Do(ctx context.Context, opts ...RequestOption) (*DeletedOrders, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not set")
	}

	r := &request{
		method:   "DELETE",
		endpoint: fmt.Sprintf("/trading/orders/%d", *s.pair_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data DeletedOrders `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *CurrencyPairOpenOrdersDeleteService) CurrencyPairId(pair_id int) *CurrencyPairOpenOrdersDeleteService {
	s.pair_id = &pair_id
	return s
}

type CreateOrderService struct {
	c *Client

	pair_id *int

	order_type    *OrderType //(BUY / SELL / STOP_LIMIT_BUY / STOP_LIMIT_SELL)
	amount        *float64
	price         *float64
	trigger_price *float64
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (*OrderInfo, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not set")
	}

	r := &request{
		method:   "POST",
		endpoint: fmt.Sprintf("/trading/orders/%d", *s.pair_id),
		secType:  secTypeAPIKey,
	}

	if s.order_type != nil {
		r.setFormParam("order_type", *s.order_type)
	} else {
		return nil, fmt.Errorf("order_type not init")
	}

	if s.amount != nil {
		r.setFormParam("amount", *s.amount)
	} else {
		return nil, fmt.Errorf("amount not init")
	}

	if s.price != nil {
		r.setFormParam("price", *s.price)
	} else {
		return nil, fmt.Errorf("price not init")
	}

	if s.trigger_price != nil {
		r.setFormParam("trigger_price", *s.trigger_price)
	} else {
		if *s.order_type == OrderType_STOP_LIMIT_BUY || *s.order_type == OrderType_STOP_LIMIT_SELL {
			return nil, fmt.Errorf("order_type not init")
		}
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data OrderInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *CreateOrderService) OrderType(order_type OrderType) *CreateOrderService {
	s.order_type = &order_type
	return s
}

func (s *CreateOrderService) Amount(amount float64) *CreateOrderService {
	s.amount = &amount
	return s
}

func (s *CreateOrderService) Price(price float64) *CreateOrderService {
	s.price = &price
	return s
}

func (s *CreateOrderService) TriggerPrice(price float64) *CreateOrderService {
	s.trigger_price = &price
	return s
}

type OrderInfoService struct {
	c *Client

	order_id *int64
}

// Do send request
func (s *OrderInfoService) Do(ctx context.Context, opts ...RequestOption) (*OrderInfo, error) {
	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/trading/order/%d", s.order_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data OrderInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *OrderInfoService) OrderId(id int64) *OrderInfoService {
	s.order_id = &id
	return s
}

type OrderDeleteService struct {
	c *Client

	order_id *int64
}

// Do send request
func (s *OrderDeleteService) Do(ctx context.Context, opts ...RequestOption) (*DeletedOrders, error) {
	r := &request{
		method:   "DELETE",
		endpoint: fmt.Sprintf("/trading/order/%d", s.order_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data DeletedOrders `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *OrderDeleteService) OrderId(id int64) *OrderDeleteService {
	s.order_id = &id
	return s
}
