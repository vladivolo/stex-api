package stex

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	ws "github.com/vladivolo/golang-socketio"
	"github.com/vladivolo/golang-socketio/transport"
)

// NOTE: All Callback func Sync!!!

type SideType string

const ()

type WssClient struct {
	APIKey    string
	BaseURL   string
	UserAgent string
	Debug     bool
	Logger    *log.Logger

	Channel   string
	EventName string
	Auth      bool

	OnMessage    interface{}
	OnDisconnect func(string)
	OnError      func(string)
	OnConnection func(string)

	c *ws.Client
}

func newWssClient(APIKey string) *WssClient {
	return &WssClient{
		APIKey:    APIKey,
		BaseURL:   "wss://socket.stex.com/socket.io/?EIO=3&transport=websocket",
		UserAgent: "Stex/golang",
		Debug:     true,
		Logger:    log.New(os.Stderr, "Stex-golang-wss ", log.LstdFlags),
	}
}

func (w *WssClient) debug(format string, v ...interface{}) {
	if w.Debug {
		w.Logger.Printf(format, v...)
	}
}

func (w *WssClient) subscribe() error {
	auth := ""
	if w.Auth == true {
		auth = w.APIKey
	}

	return w.c.Emit("subscribe", map[string]interface{}{
		"channel": w.Channel,
		"auth":    auth,
	})
}

func (w *WssClient) Do(ctx context.Context, opts ...RequestOption) error {
	if w.OnMessage == nil {
		return fmt.Errorf("Don't init onMessage callback")
	}

	var err error

	w.c, err = ws.Dial(
		w.BaseURL,
		&transport.WebsocketTransport{
			PingInterval:   15 * time.Second,
			PingTimeout:    60 * time.Second,
			ReceiveTimeout: 60 * time.Second,
			SendTimeout:    60 * time.Second,
			BufferSize:     1024 * 32,
		},
	)

	if err != nil {
		w.debug("channel %s Dial: %s", w.Channel, err)
		return err
	}

	err = w.c.On(ws.OnDisconnection, func(h *ws.Channel) {
		w.debug("channel %s OnDisconnection", w.Channel)
		if w.OnDisconnect != nil {
			w.OnDisconnect(w.Channel)
		}
	})
	if err != nil {
		return err
	}

	err = w.c.On(ws.OnError, func(h *ws.Channel) {
		w.debug("channel %s OnError", w.Channel)
		if w.OnError != nil {
			w.OnError(w.Channel)
		}
	})
	if err != nil {
		return err
	}

	err = w.c.On(ws.OnConnection, func(h *ws.Channel) {
		w.debug("channel %s OnConnection", w.Channel)
		err := w.subscribe()
		if err != nil {
			w.debug("channel %s subscribe error", w.Channel)
			if w.OnError != nil {
				w.OnError(w.Channel)
			}
		}
		if w.OnConnection != nil {
			w.OnConnection(w.Channel)
		}
	})
	if err != nil {
		return err
	}

	err = w.c.On(w.EventName, w.OnMessage)
	if err != nil {
		return err
	}

	go func() {
		select {
		case <-ctx.Done():
			w.debug("Context Done()")
			w.c.Close()
			return
		}
	}()

	return nil
}

func NewWebsocketRateChannelService() *WebsocketRateChannelService {
	return &WebsocketRateChannelService{c: newWssClient("")}
}

func NewWebsocketGlassRowChangedService() *WebsocketGlassRowChangedService {
	return &WebsocketGlassRowChangedService{c: newWssClient("")}
}
