package stex

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	ws "github.com/vladivolo/golang-socketio"
	"github.com/vladivolo/golang-socketio/transport"
)

type WssClient struct {
	sync.Mutex

	APIKey    string
	BaseURL   string
	UserAgent string
	Debug     bool
	Logger    *log.Logger

	connected bool

	onDisconnect func()
	onError      func()
	onConnection func()

	c *ws.Client
}

func NewWssClient(APIKey string) *WssClient {
	return &WssClient{
		APIKey:    APIKey,
		BaseURL:   "wss://socket.stex.com/socket.io/?EIO=3&transport=websocket",
		UserAgent: "Stex/golang",
		Debug:     false,
		Logger:    log.New(os.Stderr, "Stex-golang-wss ", log.LstdFlags),
	}
}

func (w *WssClient) debug(format string, v ...interface{}) {
	if w.Debug {
		w.Logger.Printf(format, v...)
	}
}

func (w *WssClient) Subscribe(channel string, auth bool) error {
	auth_token := map[string]interface{}{}
	if auth == true {
		auth_token = map[string]interface{}{
			"headers": map[string]interface{}{
				"Authorization": "Bearer " + w.APIKey,
			},
		}
	}

	return w.c.Emit("subscribe", map[string]interface{}{
		"channel": channel,
		"auth":    auth_token,
	})
}

func (w *WssClient) C() *ws.Client {
	return w.c
}

func (w *WssClient) SetConnected(status bool) {
	w.Lock()
	defer w.Unlock()

	w.connected = status
}

func (w *WssClient) IsConnected() bool {
	w.Lock()
	defer w.Unlock()

	return w.connected
}

func (w *WssClient) Do(ctx context.Context, opts ...RequestOption) (*WssClient, error) {
	w.Lock()
	defer w.Unlock()

	var err error

	w.c, err = ws.Dial(
		w.BaseURL,
		&transport.WebsocketTransport{
			PingInterval:   10 * time.Second,
			PingTimeout:    60 * time.Second,
			ReceiveTimeout: 60 * time.Second,
			SendTimeout:    60 * time.Second,
			BufferSize:     1024 * 32,
		},
	)

	if err != nil {
		w.debug("Dial: %s", err)
		return nil, err
	}

	err = w.c.On(ws.OnDisconnection, func(h *ws.Channel) {
		w.debug("OnDisconnection")

		w.SetConnected(false)

		if w.onDisconnect != nil {
			go w.onDisconnect()
		}
	})
	if err != nil {
		return nil, err
	}

	err = w.c.On(ws.OnError, func(h *ws.Channel) {
		w.debug("OnError:", err)
		if w.onError != nil {
			w.onError()
		}
	})
	if err != nil {
		return nil, err
	}

	err = w.c.On(ws.OnConnection, func(h *ws.Channel) {
		w.debug("OnConnection")

		w.SetConnected(true)

		if w.onConnection != nil {
			go w.onConnection()
		}
	})
	if err != nil {
		return nil, err
	}

	go func() {
		select {
		case <-ctx.Done():
			w.debug("Context Done()")
			w.c.Close()
			return
		}
	}()

	return w, nil
}

func (w *WssClient) OnConnection(f func()) *WssClient {
	w.onConnection = f
	return w
}

func (w *WssClient) OnDisconnect(f func()) *WssClient {
	w.onDisconnect = f
	return w
}

func (w *WssClient) OnError(f func()) *WssClient {
	w.onError = f
	return w
}

func NewWebsocketRateChannelService(c *WssClient) *WebsocketRateChannelService {
	return &WebsocketRateChannelService{c: c}
}

func NewWebsocketGlassRowChangedService(c *WssClient) *WebsocketGlassRowChangedService {
	return &WebsocketGlassRowChangedService{c: c}
}

func NewWebsocketUserOrderFillChannelService(c *WssClient) *WebsocketUserOrderFillChannelService {
	return &WebsocketUserOrderFillChannelService{c: c}
}

func NewWebsocketUserOrderDeletedChannelService(c *WssClient) *WebsocketUserOrderDeletedChannelService {
	return &WebsocketUserOrderDeletedChannelService{c: c}
}

func NewWebsocketUserOrderUpdateChannelService(c *WssClient) *WebsocketUserOrderUpdateChannelService {
	return &WebsocketUserOrderUpdateChannelService{c: c}
}

func NewWebsocketUserBalanceUpdateChannelService(c *WssClient) *WebsocketUserBalanceUpdateChannelService {
	return &WebsocketUserBalanceUpdateChannelService{c: c}
}
