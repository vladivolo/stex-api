package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type CurrencyPair struct {
	Id                int    `json:"id"`
	CurrencyId        int    `json:"currency_id"`
	CurrencyCode      string `json:"currency_code"`
	CurrencyName      string `json:"currency_name"`
	MarketCurrencyId  int    `json:"market_currency_id"`
	MarketCode        string `json:"market_code"`
	MarketName        string `json:"market_name"`
	MinOrderAmount    string `json:"min_order_amount"`
	MinBuyPrice       string `json:"min_buy_price"`
	MinSellPrice      string `json:"min_sell_price"`
	BuyFeePercent     string `json:"buy_fee_percent"`
	SellFeePercent    string `json:"sell_fee_percent"`
	Active            bool   `json:"active"`
	Delisted          bool   `json:"delisted"`
	PairMessage       string `json:"pair_message"`
	CurrencyPrecision int    `json:"currency_precision"`
	MarketPrecision   int    `json:"market_precision"`
	Symbol            string `json:"symbol"`
	GroupName         string `json:"group_name"`
	GroupId           int    `json:"group_id"`
	AmountMultiplier  int    `json:"amount_multiplier"`
}

type CurrencyPairsMarketListService struct {
	c             *Client
	market_symbol *string //BTC/ETH/...
}

// Do send request
func (s *CurrencyPairsMarketListService) Do(ctx context.Context, opts ...RequestOption) ([]CurrencyPair, error) {
	if s.market_symbol == nil {
		all := "ALL"

		s.market_symbol = &all
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/currency_pairs/list/%s", *s.market_symbol),
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []CurrencyPair `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *CurrencyPairsMarketListService) Market(market_symbol string) *CurrencyPairsMarketListService {
	s.market_symbol = &market_symbol
	return s
}

type CurrencyPairsGroupsService struct {
	c *Client

	group_id *int
}

// Do send request
func (s *CurrencyPairsGroupsService) Do(ctx context.Context, opts ...RequestOption) ([]CurrencyPair, error) {
	if s.group_id == nil {
		return nil, fmt.Errorf("group_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/currency_pairs/group/%d", *s.group_id),
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []CurrencyPair `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *CurrencyPairsGroupsService) GroupId(group_id int) *CurrencyPairsGroupsService {
	s.group_id = &group_id
	return s
}

type CurrencyPairInfoService struct {
	c *Client

	pair_id *int
}

// Do send request
func (s *CurrencyPairInfoService) Do(ctx context.Context, opts ...RequestOption) (*CurrencyPair, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/currency_pairs/%d", *s.pair_id),
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data CurrencyPair `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *CurrencyPairInfoService) PairId(pair_id int) *CurrencyPairInfoService {
	s.pair_id = &pair_id
	return s
}
