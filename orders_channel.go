package stex

import (
	"context"
	"fmt"
	"strings"

	ws "github.com/vladivolo/golang-socketio"
)

type TradeOrder struct {
	UserId         int64     `json:"user_id"`
	CurrencyPairId int       `json:"currency_pair_id"`
	Price          string    `json:"price"`
	Amount         string    `json:"amount"`
	Amount2        float64   `json:"amount2"`
	Date           string    `json:"date"`
	OrderType      OrderType `json:"order_type"`
}

type DeleteOrder struct {
	Id             int64  `json:"id"`
	UserId         int64  `json:"user_id"`
	CurrencyPairId int    `json:"currency_pair_id"`
	Status         string `json:"status"`
}

type UpdateOrder struct {
	Id             int64   `json:"id"`
	UserId         int64   `json:"user_id"`
	CurrencyPairId int     `json:"currency_pair_id"`
	Price          string  `json:"price"`
	Amount         string  `json:"amount"`
	Amount2        float64 `json:"amount2"`
}

type WebsocketUserOrderFillChannelService struct {
	c *WssClient

	user_id          *int64
	currency_pair_id *int

	f func(string, TradeOrder)
}

func (s *WebsocketUserOrderFillChannelService) Do(ctx context.Context, opts ...RequestOption) error {
	if s.user_id == nil {
		return fmt.Errorf("user_id not init")
	}

	if s.currency_pair_id == nil {
		return fmt.Errorf("currency_pair_id not init")
	}

	channel := fmt.Sprintf("private-trade_u%dc%d", *s.user_id, *s.currency_pair_id)

	err := s.c.Subscribe(channel, true)
	if err != nil {
		return err
	}

	err = s.c.C().On("App\\\\Events\\\\UserOrderFillCreated", func(h *ws.Channel, msg TradeOrder) {
		s.f("private-trade", msg)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *WebsocketUserOrderFillChannelService) UserId(user_id int64) *WebsocketUserOrderFillChannelService {
	s.user_id = &user_id
	return s
}

func (s *WebsocketUserOrderFillChannelService) CurrencyPairId(currency_pair_id int) *WebsocketUserOrderFillChannelService {
	s.currency_pair_id = &currency_pair_id
	return s
}

func (s *WebsocketUserOrderFillChannelService) OnMessage(f func(string, TradeOrder)) *WebsocketUserOrderFillChannelService {
	s.f = f
	return s
}

/****************************************************************************************************************************************************/

type WebsocketUserOrderDeletedChannelService struct {
	c *WssClient

	user_id          *int64
	currency_pair_id *int

	f func(string, DeleteOrder)
}

func (s *WebsocketUserOrderDeletedChannelService) Do(ctx context.Context, opts ...RequestOption) error {
	if s.user_id == nil {
		return fmt.Errorf("user_id not init")
	}

	if s.currency_pair_id == nil {
		return fmt.Errorf("currency_pair_id not init")
	}

	channel := fmt.Sprintf("private-del_order_u%dc%d", *s.user_id, *s.currency_pair_id)

	err := s.c.Subscribe(channel, true)
	if err != nil {
		return err
	}

	err = s.c.C().On("App\\\\Events\\\\UserOrderDeleted", func(h *ws.Channel, msg DeleteOrder) {
		s.f("private-delete", msg)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *WebsocketUserOrderDeletedChannelService) UserId(user_id int64) *WebsocketUserOrderDeletedChannelService {
	s.user_id = &user_id
	return s
}

func (s *WebsocketUserOrderDeletedChannelService) CurrencyPairId(currency_pair_id int) *WebsocketUserOrderDeletedChannelService {
	s.currency_pair_id = &currency_pair_id
	return s
}

func (s *WebsocketUserOrderDeletedChannelService) OnMessage(f func(string, DeleteOrder)) *WebsocketUserOrderDeletedChannelService {
	s.f = f
	return s
}

/****************************************************************************************************************************************************/

type WebsocketUserOrderUpdateChannelService struct {
	c *WssClient

	user_id          *int64
	currency_pair_id *int
	order_type       *OrderType

	f func(string, UpdateOrder)
}

func (s *WebsocketUserOrderUpdateChannelService) Do(ctx context.Context, opts ...RequestOption) error {
	if s.user_id == nil {
		return fmt.Errorf("user_id not init")
	}

	if s.order_type == nil {
		return fmt.Errorf("order_type not init")
	}

	if s.currency_pair_id == nil {
		return fmt.Errorf("currency_pair_id not init")
	}

	channel := fmt.Sprintf("private-%s_user_data_u%dc%d", *s.order_type, *s.user_id, *s.currency_pair_id)
	err := s.c.Subscribe(channel, true)
	if err != nil {
		return err
	}

	err = s.c.C().On("App\\\\Events\\\\UserOrder", func(h *ws.Channel, msg UpdateOrder) {
		s.f(string(*s.order_type), msg)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *WebsocketUserOrderUpdateChannelService) UserId(user_id int64) *WebsocketUserOrderUpdateChannelService {
	s.user_id = &user_id
	return s
}

func (s *WebsocketUserOrderUpdateChannelService) CurrencyPairId(currency_pair_id int) *WebsocketUserOrderUpdateChannelService {
	s.currency_pair_id = &currency_pair_id
	return s
}

func (s *WebsocketUserOrderUpdateChannelService) OrderType(order_type OrderType) *WebsocketUserOrderUpdateChannelService {
	order_type = OrderType(strings.ToLower(string(order_type)))
	s.order_type = &order_type
	return s
}

func (s *WebsocketUserOrderUpdateChannelService) OnMessage(f func(string, UpdateOrder)) *WebsocketUserOrderUpdateChannelService {
	s.f = f
	return s
}
