package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type CurrencyPairTicker struct {
	Id               int                `json:"id"`
	AmountMultiplier int                `json:"amount_multiplier"`
	CurrencyCode     string             `json:"currency_code"`
	MarketCode       string             `json:"market_code"`
	CurrencyName     string             `json:"currency_name"`
	MarketName       string             `json:"market_name"`
	Symbol           string             `json:"symbol"`
	GroupName        string             `json:"group_name"`
	GroupId          int                `json:"group_id"`
	Ask              string             `json:"ask"`
	Bid              string             `json:"bid"`
	Last             string             `json:"last"`
	Low              string             `json:"low"`
	High             string             `json:"high"`
	Open             string             `json:"open"`
	Volume           string             `json:"volume"`
	VolumeQuote      string             `json:"volumeQuote"`
	FiatsRate        map[string]float64 `json:"fiatsRate"`
	Timestamp        int64              `json:"timestamp"`
}

type CurrencyPairsTickerService struct {
	c *Client
}

// Do send request
func (s *CurrencyPairsTickerService) Do(ctx context.Context, opts ...RequestOption) ([]CurrencyPairTicker, error) {
	r := &request{
		method:   "GET",
		endpoint: "/public/ticker",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []CurrencyPairTicker `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

type CurrencyPairTickerService struct {
	c *Client

	pair_id *int
}

// Do send request
func (s *CurrencyPairTickerService) Do(ctx context.Context, opts ...RequestOption) (*CurrencyPairTicker, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/ticker/%d", *s.pair_id),
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data CurrencyPairTicker `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *CurrencyPairTickerService) CurrencyPairId(pair_id int) *CurrencyPairTickerService {
	s.pair_id = &pair_id
	return s
}
