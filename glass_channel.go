package stex

import (
	"fmt"
	"strings"

	ws "github.com/vladivolo/golang-socketio"
)

type WebsocketGlassRowChangedService struct {
	c *WssClient

	trade_type       *TradeType
	currency_pair_id *int

	f func(TradeType, Order)
}

func (s *WebsocketGlassRowChangedService) Do() error {
	if s.trade_type == nil {
		return fmt.Errorf("trade_type not init")
	}

	if s.currency_pair_id == nil {
		return fmt.Errorf("currency_pair_id not init")
	}

	channel := fmt.Sprintf("%s_data%d", strings.ToLower(string(*s.trade_type)), *s.currency_pair_id)
	err := s.c.Subscribe(channel, false)
	if err != nil {
		return err
	}

	err = s.c.C().On("App\\\\Events\\\\GlassRowChanged", func(h *ws.Channel, msg Order) {
		s.f(*s.trade_type, msg)
	}, channel)
	if err != nil {
		return err
	}

	return nil
}

func (s *WebsocketGlassRowChangedService) TradeType(side TradeType) *WebsocketGlassRowChangedService {
	s.trade_type = &side
	return s
}

func (s *WebsocketGlassRowChangedService) CurrencyPairId(currency_pair_id int) *WebsocketGlassRowChangedService {
	s.currency_pair_id = &currency_pair_id
	return s
}

func (s *WebsocketGlassRowChangedService) OnMessage(f func(TradeType, Order)) *WebsocketGlassRowChangedService {
	s.f = f
	return s
}
