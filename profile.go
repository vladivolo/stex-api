package stex

import (
	"context"
	"encoding/json"
	"fmt"
)

type ProfileVerifications struct {
	Cryptonomica bool `json:"cryptonomica"`
	Privatbank   bool `json:"privatbank"`
	Stex         bool `json:"stex"`
}

type TradingFeeLevel struct {
	NotVerified  string `json:"not_verified"`
	Cryptonomica string `json:"cryptonomica"`
	Privatbank   string `json:"privatbank"`
	Stex         string `json:"stex"`
}

type ReferralProgram struct {
	ReferralCode string `json:"referral_code"`
	Members      int    `json:"members"`
	Invited      bool   `json:"invited"`
}

type Balance struct {
	Balance       string `json:"balance"`
	FrozenBalance string `json:"frozen_balance"`
	BonusBalance  string `json:"bonus_balance"`
	TotalBalance  string `json:"total_balance"`
}

type ProfileInfo struct {
	Email                 string               `json:"email"`
	Username              string               `json:"username"`
	UserId                int64                `json:"user_id"`
	Verifications         ProfileVerifications `json:"verifications"`
	TradingFeeLevels      TradingFeeLevel      `json:"trading_fee_levels"`
	ApiWithdrawalsAllowed bool                 `json:"api_withdrawals_allowed"`
	ReferralProgram       ReferralProgram      `json:"referral_program"`
	ApproxBalance         map[string]Balance   `json:"approx_balance"`
}

type Wallet struct {
	Id              int64              `json:"id"`
	CurrencyId      int                `json:"currency_id"`
	Delisted        bool               `json:"delisted"`
	Disabled        bool               `json:"disabled"`
	DisableDeposits bool               `json:"disable_deposits"`
	CurrencyCode    string             `json:"currency_code"`
	CurrencyName    string             `json:"currency_name"`
	OfficialUrl     string             `json:"official_url"`
	Rates           map[string]float64 `json:"rates"`
	Balance         string             `json:"balance"`
	FrozenBalance   string             `json:"frozen_balance"`
	BonusBalance    string             `json:"bonus_balance"`
}

type Address struct {
	Address                        string `json:"address"`
	AddressName                    string `json:"address_name"`
	AdditionalAddressParameter     string `json:"additional_address_parameter"`
	AdditionalAddressParameterName string `json:"additional_address_parameter_name"`
	Notification                   string `json:"notification"`
	ProtocolId                     int    `json:"protocol_id"`
	ProtocolName                   string `json:"protocol_name"`
}

type WalletAdv struct {
	Id                            int64              `json:"id"`
	CurrencyId                    int                `json:"currency_id"`
	Delisted                      bool               `json:"delisted"`
	Disabled                      bool               `json:"disabled"`
	DisableDeposits               bool               `json:"disable_deposits"`
	Code                          string             `json:"code"`
	Balance                       string             `json:"balance"`
	FrozenBalance                 string             `json:"frozen_balance"`
	BonusBalance                  string             `json:"bonus_balance"`
	DepositAddress                Address            `json:"deposit_address"`
	MultiDepositAddress           Address            `json:"multi_deposit_address"`
	WithdrawalAdditionalFieldName string             `json:"withdrawal_additional_field_name"`
	Rates                         map[string]float64 `json:"rates"`
}

type ProfileInfoService struct {
	c *Client
}

func (s *ProfileInfoService) Do(ctx context.Context, opts ...RequestOption) (*ProfileInfo, error) {
	r := &request{
		method:   "GET",
		endpoint: "/profile/info",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data ProfileInfo `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

type ProfileWalletListService struct {
	c *Client

	sort   *SortOrder
	sortBy *SortBalanceField
}

func (s *ProfileWalletListService) Do(ctx context.Context, opts ...RequestOption) ([]Wallet, error) {
	r := &request{
		method:   "GET",
		endpoint: "/profile/wallets",
		secType:  secTypeAPIKey,
	}

	if s.sort != nil {
		r.setParam("sort", *s.sort)
	}

	if s.sortBy != nil {
		r.setParam("sortBy", *s.sortBy)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data []Wallet `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, err
}

func (s *ProfileWalletListService) Order(order SortOrder) *ProfileWalletListService {
	s.sort = &order
	return s
}

func (s *ProfileWalletListService) SortBy(field SortBalanceField) *ProfileWalletListService {
	s.sortBy = &field
	return s
}

type ProfileWalletInfoService struct {
	c *Client

	wallet_id *int64
}

func (s *ProfileWalletInfoService) Do(ctx context.Context, opts ...RequestOption) (*WalletAdv, error) {
	if s.wallet_id == nil {
		return nil, fmt.Errorf("wallet_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/profile/wallets/%d", *s.wallet_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data WalletAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileWalletInfoService) WalletId(id int64) *ProfileWalletInfoService {
	s.wallet_id = &id
	return s
}

type ProfileWalletCreateService struct {
	c *Client

	currency_id *int64
	protocol_id *int
}

func (s *ProfileWalletCreateService) Do(ctx context.Context, opts ...RequestOption) (*WalletAdv, error) {
	if s.currency_id == nil {
		return nil, fmt.Errorf("currency_id not init")
	}

	r := &request{
		method:   "POST",
		endpoint: fmt.Sprintf("/profile/wallets/%d", *s.currency_id),
		secType:  secTypeAPIKey,
	}

	if s.protocol_id != nil {
		r.setParam("protocol_id", *s.protocol_id)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data WalletAdv `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileWalletCreateService) CurrencyId(id int64) *ProfileWalletCreateService {
	s.currency_id = &id
	return s
}

func (s *ProfileWalletCreateService) ProtocolId(id int) *ProfileWalletCreateService {
	s.protocol_id = &id
	return s
}

type ProfileWalletAddressInfoService struct {
	c *Client

	wallet_id   *int64
	protocol_id *int
}

func (s *ProfileWalletAddressInfoService) Do(ctx context.Context, opts ...RequestOption) (*Address, error) {
	if s.wallet_id == nil {
		return nil, fmt.Errorf("wallet_id not init")
	}

	r := &request{
		method:   "GET",
		endpoint: fmt.Sprintf("/profile/wallets/address/%d", *s.wallet_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	if s.protocol_id != nil {
		r.setParam("protocol_id", *s.protocol_id)
	}

	res := struct {
		APIError
		Data Address `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileWalletAddressInfoService) WalletId(id int64) *ProfileWalletAddressInfoService {
	s.wallet_id = &id
	return s
}

func (s *ProfileWalletAddressInfoService) ProtocolId(id int) *ProfileWalletAddressInfoService {
	s.protocol_id = &id
	return s
}

type ProfileWalletAddressCreateService struct {
	c *Client

	wallet_id   *int64
	protocol_id *int
}

func (s *ProfileWalletAddressCreateService) Do(ctx context.Context, opts ...RequestOption) (*Address, error) {
	if s.wallet_id == nil {
		return nil, fmt.Errorf("wallet_id not init")
	}

	r := &request{
		method:   "POST",
		endpoint: fmt.Sprintf("/profile/wallets/address/%d", *s.wallet_id),
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	if s.protocol_id != nil {
		r.setParam("protocol_id", *s.protocol_id)
	}

	res := struct {
		APIError
		Data Address `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileWalletAddressCreateService) WalletId(id int64) *ProfileWalletAddressCreateService {
	s.wallet_id = &id
	return s
}

func (s *ProfileWalletAddressCreateService) ProtocolId(id int) *ProfileWalletAddressCreateService {
	s.protocol_id = &id
	return s
}
