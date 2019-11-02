package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type Order struct {
	CurrencyPairId   int     `json:"currency_pair_id"`
	Amount           string  `json:"amount"`
	Price            string  `json:"price"`
	Amount2          string  `json:"amount2"`
	Count            int     `json:"count"`
	CumulativeAmount float64 `json:"cumulative_amount"`
}

type OrderBook struct {
	Ask            []Order `json:"ask"`
	Bid            []Order `json:"bid"`
	AskTotalAmount float64 `json:"ask_total_amount"`
	BidTotalAmount float64 `json:"bid_total_amount"`
}

type CurrencyPairOrderbookService struct {
	c *Client

	pair_id    *int
	limit_bids *int
	limit_asks *int
}

// Do send request
func (s *CurrencyPairOrderbookService) Do(ctx context.Context, opts ...RequestOption) (*OrderBook, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/orderbook/%d", *s.pair_id),
		secType:  secTypeNone,
	}

	if s.limit_bids != nil {
		r.setParam("limit_bids", *s.limit_bids)
	}

	if s.limit_asks != nil {
		r.setParam("limit_asks", *s.limit_asks)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data OrderBook `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *CurrencyPairOrderbookService) BidsLimit(limit int) *CurrencyPairOrderbookService {
	s.limit_bids = &limit
	return s
}

func (s *CurrencyPairOrderbookService) AsksLimit(limit int) *CurrencyPairOrderbookService {
	s.limit_asks = &limit
	return s
}
