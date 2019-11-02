package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type Candle struct {
	Time   int64   `json:"time"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Volume float64 `json:"volume"`
}

type CurrencyPairChartService struct {
	c *Client

	// A currency pair ID you want to get candles for
	pair_id *int
	// Candle size 1 stays for 1 minute, 5 - 5 minutes and so on. 1D - stays for 1 day
	candle_type *CandleType
	// Timestamp in second. Should be less then timeEnd
	tm_start *int64
	// Timestamp in second. Should be greater then timeStart
	tm_end *int64

	limit  *int
	offset *int
}

// Do send request
func (s *CurrencyPairChartService) Do(ctx context.Context, opts ...RequestOption) ([]Candle, error) {
	if s.pair_id == nil {
		return nil, fmt.Errorf("pair_id not init")
	}

	if s.candle_type == nil {
		return nil, fmt.Errorf("candle_type not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/chart/%d/%s", *s.pair_id, *s.candle_type),
		secType:  secTypeNone,
	}

	if s.tm_start != nil {
		r.setParam("timeStart", *s.tm_start)
	} else {
		return nil, fmt.Errorf("TmStart not init")
	}

	if s.tm_end != nil {
		r.setParam("timeEnd", *s.tm_end)
	} else {
		return nil, fmt.Errorf("TmEnd not init")
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
		Data []Candle `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *CurrencyPairChartService) CurrencyPairId(pair_id int) *CurrencyPairChartService {
	s.pair_id = &pair_id
	return s
}

func (s *CurrencyPairChartService) CandleType(candle_type CandleType) *CurrencyPairChartService {
	s.candle_type = &candle_type
	return s
}

func (s *CurrencyPairChartService) Limit(limit int) *CurrencyPairChartService {
	s.limit = &limit
	return s
}

func (s *CurrencyPairChartService) Offset(offset int) *CurrencyPairChartService {
	s.offset = &offset
	return s
}
