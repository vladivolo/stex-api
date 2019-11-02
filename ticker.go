package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

/*
"data": [
      {
        "id": 1,
        "amount_multiplier": 1,
        "currency_code": "ETH",
        "market_code": "BTC",
        "currency_name": "Etherium",
        "market_name": "Bitcoin",
        "symbol": "ETH_BTC",
        "group_name": "FIAT coins",
        "group_id": 1,
        "ask": "0.03377988",
        "bid": "0.03350001",
        "last": "0.0337",
        "low": "0.03320157",
        "high": "0.0341",
        "open": "0.03340002",
        "volume": "5.1939",
        "volumeQuote": "154.12169946",
        "fiatsRate": {
          "BTC": 0.000001
        },
        "timestamp": 1538737692
      }
*/

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

func (s *CurrencyPairTickerService) PairId(pair_id int) *CurrencyPairTickerService {
	s.pair_id = &pair_id
	return s
}
