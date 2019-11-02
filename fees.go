package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type Fees struct {
	SellFee string `json:"sell_fee"`
	BuyFee  string `json:"buy_fee"`
}

type CurrencyPairFeeService struct {
	c *Client

	pair_id *int
}

// Do send request
func (s *CurrencyPairFeeService) Do(ctx context.Context, opts ...RequestOption) (*Fees, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/trading/fees/%d", *s.pair_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data Fees `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *CurrencyPairFeeService) CurrencyPairId(pair_id int) *CurrencyPairFeeService {
	s.pair_id = &pair_id
	return s
}
