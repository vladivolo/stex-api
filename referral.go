package stex

import (
	"context"
	"encoding/json"
)

type Referral struct {
	ReferralCode string `json:"referral_code"`
	Members      int    `json:"members"`
	Invited      bool   `json:"invited"`
}

type ProfileReferralCreateService struct {
	c *Client
}

// Do send request
func (s *ProfileReferralCreateService) Do(ctx context.Context, opts ...RequestOption) (*Referral, error) {
	r := &request{
		method:   "POST",
		endpoint: "/profile/referral/program",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data Referral `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

type ProfileReferralSetService struct {
	c *Client

	code *string
}

// Do send request
func (s *ProfileReferralSetService) Do(ctx context.Context, opts ...RequestOption) (*Referral, error) {
	r := &request{
		method:   "POST",
		endpoint: "/profile/referral/insert",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := struct {
		APIError
		Data Referral `json:"data"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, err
}

func (s *ProfileReferralSetService) Code(code string) *ProfileReferralSetService {
	s.code = &code
	return s
}
