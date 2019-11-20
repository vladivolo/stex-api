package stex

import (
	"fmt"

	ws "github.com/vladivolo/golang-socketio"
)

type UpdateBalance struct {
	Id            int64  `json:"id"`
	Code          string `json:"currency_code"`
	Balance       string `json:"balance"`
	FrozenBalance string `json:"frozen_balance"`
	BonusBalance  string `json:"bonus_balance"`
	TotalBalance  string `json:"total_balance"`
}

type WebsocketUserBalanceUpdateChannelService struct {
	c *WssClient

	wallet_id *int64

	f func(string, UpdateBalance)
}

func (s *WebsocketUserBalanceUpdateChannelService) Do() error {
	if s.wallet_id == nil {
		return fmt.Errorf("wallet_id not init")
	}

	channel := fmt.Sprintf("private-balance_changed_w_%d", *s.wallet_id)

	err := s.c.Subscribe(channel, true)
	if err != nil {
		return err
	}

	err = s.c.C().On("App\\\\Events\\\\BalanceChanged", func(h *ws.Channel, msg UpdateBalance) {
		s.f("private-balance", msg)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *WebsocketUserBalanceUpdateChannelService) WalletId(wallet_id int64) *WebsocketUserBalanceUpdateChannelService {
	s.wallet_id = &wallet_id
	return s
}

func (s *WebsocketUserBalanceUpdateChannelService) OnMessage(f func(string, UpdateBalance)) *WebsocketUserBalanceUpdateChannelService {
	s.f = f
	return s
}
