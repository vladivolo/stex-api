package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type ProtocolSpecificSettings struct {
	ProtocolName            string  `json:"protocol_name"`
	ProtocolId              int     `json:"protocol_id"`
	Active                  bool    `json:"active"`
	WithdrawalFeeCurrencyId int     `json:"withdrawal_fee_currency_id"`
	WithdrawalFeeConst      float64 `json:"withdrawal_fee_const"`
	WithdrawalFeePercent    float64 `json:"withdrawal_fee_percent"`
	BlockExplorerUrl        string  `json:"block_explorer_url"`
}

type CurrencyInfo struct {
	Id                        int                        `json:"id"`
	Code                      string                     `json:"code"`
	Name                      string                     `json:"name"`
	Active                    bool                       `json:"active"`
	Delisted                  bool                       `json:"delisted"`
	Precision                 int                        `json:"precision"`
	MinimumWithdrawalAmount   string                     `json:"minimum_withdrawal_amount"`
	MinimumDepositAmount      string                     `json:"minimum_deposit_amount"`
	DepositFeeCurrencyId      int                        `json:"deposit_fee_currency_id"`
	DepositFeeCurrencyCode    string                     `json:"deposit_fee_currency_code"`
	DepositFeeConst           string                     `json:"deposit_fee_const"`
	DepositFeePercent         string                     `json:"deposit_fee_percent"`
	WithdrawalFeeCurrencyId   int                        `json:"withdrawal_fee_currency_id"`
	WithdrawalFeeCurrencyCode string                     `json:"withdrawal_fee_currency_code"`
	WithdrawalFeeConst        string                     `json:"withdrawal_fee_const"`
	WithdrawalFeePercent      string                     `json:"withdrawal_fee_percent"`
	BlockExplorerUrl          string                     `json:"block_explorer_url"`
	ProtocolSpecificSettings  []ProtocolSpecificSettings `json:"protocol_specific_settings"`
}

type AvailableCurrenciesService struct {
	c *Client
}

// Do send request
func (s *AvailableCurrenciesService) Do(ctx context.Context, opts ...RequestOption) ([]CurrencyInfo, error) {
	r := &request{
		method:   "GET",
		endpoint: "/public/currencies",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []CurrencyInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

type CurrencyInfoByIdService struct {
	currency_id *int
	c           *Client
}

type CurrencyInfoByIdResponse struct {
	APIError
	Data CurrencyInfo `json:"data"`
}

// Do send request
func (s *CurrencyInfoByIdService) Do(ctx context.Context, opts ...RequestOption) (CurrencyInfo, error) {
	if s.currency_id == nil {
		return CurrencyInfo{}, fmt.Errorf("CurrencyId not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/public/currencies/%d", *s.currency_id),
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return CurrencyInfo{}, err
	}

	res := struct {
		APIError
		Data CurrencyInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return CurrencyInfo{}, err
	}

	return res.Data, err
}

func (s *CurrencyInfoByIdService) Id(currency_id int) *CurrencyInfoByIdService {
	s.currency_id = &currency_id
	return s
}
