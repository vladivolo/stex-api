package stex

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type CurrencyPairTrades struct {
	Id        int64  `json:"id"`
	Price     string `json:"price"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Timestamp int64  `json:"timestamp"`
}

type CurrencyPairTradesService struct {
	c       *Client
	pair_id int
	sort    SortOrder
	from    int64
	till    int64
	limit   int
	offset  int
}

// Do send request
func (s *CurrencyPairTradesService) Do(ctx context.Context, opts ...RequestOption) ([]CurrencyPairTrades, error) {
	if s.pair_id == 0 {
		return nil, fmt.Errorf("pair_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/trades/%d", s.pair_id),
		secType:  secTypeNone,
	}

	if s.from != 0 {
		r.setParam("from", s.from)
	}

	if s.till != 0 {
		r.setParam("till", s.till)
	}

	if s.sort != "" {
		r.setParam("sort", s.sort)
	}

	if s.limit != 0 {
		r.setParam("limit", s.limit)
	}

	r.setParam("offset", s.offset)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []CurrencyPairTrades `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *CurrencyPairTradesService) CurrencyPairId(pair_id int) *CurrencyPairTradesService {
	s.pair_id = pair_id
	return s
}

func (s *CurrencyPairTradesService) Sort(order SortOrder) *CurrencyPairTradesService {
	s.sort = order
	return s
}

func (s *CurrencyPairTradesService) From(tm time.Time) *CurrencyPairTradesService {
	s.from = tm.Unix()
	return s
}

func (s *CurrencyPairTradesService) Till(tm time.Time) *CurrencyPairTradesService {
	s.till = tm.Unix()
	return s
}

func (s *CurrencyPairTradesService) Limit(limit int) *CurrencyPairTradesService {
	s.limit = limit
	return s
}

func (s *CurrencyPairTradesService) Offset(offset int) *CurrencyPairTradesService {
	s.offset = offset
	return s
}
