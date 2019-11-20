package stex

import (
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

	f func(string, RateMessage)
}

func (s *WebsocketRateChannelService) Do() error {
	err := s.c.Subscribe("rate", false)
	if err != nil {
		return err
	}

	err = s.c.C().On("App\\\\Events\\\\Ticker", func(h *ws.Channel, msg RateMessage) {
		s.f("rate", msg)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *WebsocketRateChannelService) OnMessage(f func(string, RateMessage)) *WebsocketRateChannelService {
	s.f = f
	return s
}
