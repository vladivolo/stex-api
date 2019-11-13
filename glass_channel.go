package stex

import (
	"context"
	"fmt"
	"strings"

	ws "github.com/vladivolo/golang-socketio"
)

type WebsocketGlassRowChangedService struct {
	c *WssClient

	trade_type       *TradeType
	currency_pair_id *int

	onMessage func(string, Order)
}

func (s *WebsocketGlassRowChangedService) Do(ctx context.Context, opts ...RequestOption) error {
	s.c.Channel = fmt.Sprintf("%s_data%d", strings.ToLower(string(*s.trade_type)), *s.currency_pair_id)
	s.c.EventName = "App\\\\Events\\\\GlassRowChanged"

	return s.c.Do(ctx, opts...)
}

func (s *WebsocketGlassRowChangedService) TradeType(side TradeType) *WebsocketGlassRowChangedService {
	s.trade_type = &side
	return s
}

func (s *WebsocketGlassRowChangedService) CurrencyPairId(currency_pair_id int) *WebsocketGlassRowChangedService {
	s.currency_pair_id = &currency_pair_id
	return s
}

func (s *WebsocketGlassRowChangedService) OnMessage(f func(string, Order)) *WebsocketGlassRowChangedService {
	/*
		if s.trade_type == nil {
			s.c.Logger.Fatal(fmt.Errorf("Init trade_type before set OnMessage"))
		}

		if s.currency_pair_id == nil {
			s.c.Logger.Fatal(fmt.Errorf("Init currency_pair_id before set OnMessage"))
		}
	*/
	s.onMessage = f

	// Hack. Will be refactored after understanding the package reflect
	s.c.OnMessage = func(h *ws.Channel, msg Order) {
		s.onMessage(s.c.Channel, msg)
	}

	return s
}

func (s *WebsocketGlassRowChangedService) OnDisconnect(f func(string)) *WebsocketGlassRowChangedService {
	s.c.OnDisconnect = f
	return s
}

func (s *WebsocketGlassRowChangedService) OnError(f func(string)) *WebsocketGlassRowChangedService {
	s.c.OnError = f
	return s
}

func (s *WebsocketGlassRowChangedService) OnConnection(f func(string)) *WebsocketGlassRowChangedService {
	s.c.OnConnection = f
	return s
}
