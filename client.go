package stex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SortOrder string
type CandleType string
type OrderType string

const (
	secTypeNone secType = iota
	secTypeAPIKey

	SortDesc SortOrder = "DESC"
	SortAsc  SortOrder = "ASC"

	CandleType1m  CandleType = "1"
	CandleType5m  CandleType = "5"
	CandleType30m CandleType = "30"
	CandleType1h  CandleType = "60"
	CandleType4h  CandleType = "240"
	CandleType12h CandleType = "720"
	CandleType1d  CandleType = "1D"

	OrderType_BUY             OrderType = "BUY"
	OrderType_SELL            OrderType = "SELL"
	OrderType_STOP_LIMIT_BUY  OrderType = "STOP_LIMIT_BUY"
	OrderType_STOP_LIMIT_SELL OrderType = "STOP_LIMIT_SELL"
)

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	APIKey     string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	do         doFunc
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    "https://api3.stex.com",
		UserAgent:  "Stex/golang",
		HTTPClient: http.DefaultClient,
		Debug:      true,
		Logger:     log.New(os.Stderr, "Stex-golang ", log.LstdFlags),
	}
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}

	header.Set("accept", "application/json")

	if r.secType == secTypeAPIKey {
		header.Set("Authorization", "Bearer "+c.APIKey)
	}

	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= 400 {
		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

func (c *Client) NewAvailableCurrenciesService() *AvailableCurrenciesService {
	return &AvailableCurrenciesService{c: c}
}

func (c *Client) NewCurrencyInfoByIdService() *CurrencyInfoByIdService {
	return &CurrencyInfoByIdService{c: c}
}

func (c *Client) NewAvailableMarketsService() *AvailableMarketsService {
	return &AvailableMarketsService{c: c}
}

func (c *Client) NewPairsGroupsService() *PairsGroupsService {
	return &PairsGroupsService{c: c}
}

func (c *Client) NewCurrencyPairsListService() *CurrencyPairsListService {
	return &CurrencyPairsListService{c: c}
}

func (c *Client) NewCurrencyPairsGroupsService() *CurrencyPairsGroupsService {
	return &CurrencyPairsGroupsService{c: c}
}

func (c *Client) NewCurrencyPairInfoService() *CurrencyPairInfoService {
	return &CurrencyPairInfoService{c: c}
}

func (c *Client) NewCurrencyPairsTickerService() *CurrencyPairsTickerService {
	return &CurrencyPairsTickerService{c: c}
}

func (c *Client) NewCurrencyPairTickerService() *CurrencyPairTickerService {
	return &CurrencyPairTickerService{c: c}
}

func (c *Client) NewCurrencyPairTradesService() *CurrencyPairTradesService {
	return &CurrencyPairTradesService{c: c}
}

func (c *Client) NewCurrencyPairOrderbookService() *CurrencyPairOrderbookService {
	return &CurrencyPairOrderbookService{c: c}
}

func (c *Client) NewCurrencyPairChartService() *CurrencyPairChartService {
	return &CurrencyPairChartService{c: c}
}

func (c *Client) NewDepositStatusesService() *DepositStatusesService {
	return &DepositStatusesService{c: c}
}

func (c *Client) NewDepositStatusByIdService() *DepositStatusByIdService {
	return &DepositStatusByIdService{c: c}
}

func (c *Client) NewWithdrawalStatusesService() *WithdrawalStatusesService {
	return &WithdrawalStatusesService{c: c}
}

func (c *Client) NewWithdrawalStatusByIdService() *WithdrawalStatusByIdService {
	return &WithdrawalStatusByIdService{c: c}
}

func (c *Client) NewCurrencyPairFeeService() *CurrencyPairFeeService {
	return &CurrencyPairFeeService{c: c}
}

func (c *Client) NewOpenOrdersListService() *OpenOrdersListService {
	return &OpenOrdersListService{c: c}
}

func (c *Client) NewOpenOrdersDeleteService() *OpenOrdersDeleteService {
	return &OpenOrdersDeleteService{c: c}
}

func (c *Client) NewCurrencyPairOpenOrdersListService() *CurrencyPairOpenOrdersListService {
	return &CurrencyPairOpenOrdersListService{c: c}
}

func (c *Client) NewCurrencyPairOpenOrdersDeleteService() *CurrencyPairOpenOrdersDeleteService {
	return &CurrencyPairOpenOrdersDeleteService{c: c}
}

func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{c: c}
}

func (c *Client) NewOrderInfoService() *OrderInfoService {
	return &OrderInfoService{c: c}
}

func (c *Client) NewOrderDeleteService() *OrderDeleteService {
	return &OrderDeleteService{c: c}
}
