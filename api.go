// Package allyinvest wraps the Ally Invest API.
package allyinvest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mrjones/oauth"
)

// API is the Ally Invest API.
type API interface {
	GetAccounts() (GetAccountsOutput, error)
	GetAccountsBalances() (GetAccountsBalancesOutput, error)
	GetAccount(GetAccountInput) (GetAccountOutput, error)
	GetAccountBalances(GetAccountBalancesInput) (GetAccountBalancesOutput, error)
	GetAccountHistory(GetAccountHistoryInput) (GetAccountHistoryOutput, error)
	GetAccountHoldings(GetAccountHoldingsInput) (GetAccountHoldingsOutput, error)

	GetOrders()
	CreateOrder()
	PreviewOrder()

	GetMarketClock()
	GetMarketNews()
	SearchMarketNews()
	GetMarketTimesales()
	GetMarketToplists()

	GetQuotes()

	SearchOptions(SearchOptionsInput) (SearchOptionsOutput, error)
	GetOptionsStrikes()
	GetOptionsExpirations()

	GetMemberProfile()

	GetUtilityStatus()
	GetUtilityVersion()

	GetWatchlists()
	CreateWatchlist()
	GetWatchlist()
	DeleteWatchlist()
	AddSymbolToWatchlist()
	DeleteSymbolFromWatchlist()
}

const (
	requestTokenUrl   = "https://developers.tradeking.com/oauth/request_token"
	authorizeTokenUrl = "https://developers.tradeking.com/oauth/authorize"
	accessTokenUrl    = "https://developers.tradeking.com/oauth/access_token"

	allyinvestV1 = "https://devapi.invest.ally.com/v1"
)

// New returns a new API for use.
func New(k ApplicationKeys) API {
	a := &api{}

	c := oauth.NewConsumer(
		k.ConsumerKey,
		k.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenUrl,
			AuthorizeTokenUrl: authorizeTokenUrl,
			AccessTokenUrl:    accessTokenUrl,
		},
	)

	a.client, _ = c.MakeHttpClient(
		&oauth.AccessToken{
			Token:  k.AccessToken,
			Secret: k.AccessSecret,
		},
	)

	return a
}

// NewWithClient creates a new API with a custom HTTPClient
func NewWithClient(client HTTPClient) API {
	return &api{
		client: client,
	}
}

// ApplicationKeys are the OAuth keys associated with a given account.
type ApplicationKeys struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// Api privately implements API.
type api struct {
	client HTTPClient
	API
}

// HTTPClient defines the HTTP client to be used by the application.
// The stdlib's *http.Client implements this interface.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Query indicates that the provided object can be encoded into
// an URL query string.
type query interface {
	Encode() string
}

func (a *api) httpGet(route string, query query, output interface{}) error {
	route = strings.TrimPrefix(route, "/")
	u := fmt.Sprintf("%s/%s.json", allyinvestV1, route)
	if q := query.Encode(); q != "" {
		u = fmt.Sprintf("%s?%s", u, q)
	}

	req, _ := http.NewRequest(http.MethodGet, u, nil)
	fmt.Println(req.URL)
	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, &output)
}
