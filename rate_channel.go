package stex

import (
	"context"

	ws "github.com/vladivolo/golang-socketio"
)

type RateMessage struct {
	Id              int    `json:"id"`
	ClosedOrders    int    `json:"closedOrders"`
	LastPriceDayAgo string `json:"lastPriceDayAgo"`
	MaxBuy          string `json:"maxBuy"`
	MinSell         string `json:"minSell"`
	VolumeSum       string `json:"volumeSum"`
	MarketVolume    string `json:"market_volume"`
	LastPrice       string `json:"lastPrice"`
	Spread          string `json:"spread"`
	Precision       int    `json:"precision"`
}

type WebsocketRateChannelService struct {
	c *WssClient

	onMessage func(string, RateMessage)
}

func (s *WebsocketRateChannelService) Do(ctx context.Context, opts ...RequestOption) error {
	s.c.Channel = "rate"
	s.c.EventName = "App\\\\Events\\\\Ticker"

	return s.c.Do(ctx, opts...)
}

func (s *WebsocketRateChannelService) OnMessage(f func(string, RateMessage)) *WebsocketRateChannelService {
	s.onMessage = f

	// Hack. Will be refactored after understanding the package reflect
	s.c.OnMessage = func(h *ws.Channel, msg RateMessage) {
		s.onMessage(s.c.Channel, msg)
	}

	return s
}

func (s *WebsocketRateChannelService) OnDisconnect(f func(string)) *WebsocketRateChannelService {
	s.c.OnDisconnect = f
	return s
}

func (s *WebsocketRateChannelService) OnError(f func(string)) *WebsocketRateChannelService {
	s.c.OnError = f
	return s
}

func (s *WebsocketRateChannelService) OnConnection(f func(string)) *WebsocketRateChannelService {
	s.c.OnConnection = f
	return s
}
