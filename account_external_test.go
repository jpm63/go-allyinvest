package allyinvest_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/jpm63/go-allyinvest"
)

type mockHTTPClient struct {
	DoValue *http.Response
	DoError error
}

func (m mockHTTPClient) Do(*http.Request) (*http.Response, error) {
	return m.DoValue, m.DoError
}

func TestGetAccountHoldings(t *testing.T) {
	t.Run("InvalidInput", func(t *testing.T) {
		a := allyinvest.NewWithClient(http.DefaultClient)
		_, err := a.GetAccountHoldings(allyinvest.GetAccountHoldingsInput{})
		if err == nil {
			t.Errorf("got nil, want err")
		}
	})

	t.Run("Success", func(t *testing.T) {
		a := allyinvest.NewWithClient(mockHTTPClient{
			DoValue: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte{})),
			},
		})

		_, err := a.GetAccountHoldings(allyinvest.GetAccountHoldingsInput{
			AccountID: allyinvest.String("s"),
		})

		if err == nil {
			t.Errorf("got nil, want err")
		}
	})
}

func TestGetAccountHistoryInputEncode(t *testing.T) {
	r := allyinvest.AccountHistoryDateRangeToday
	d := allyinvest.AccountHistoryTransactionTrade
	i := allyinvest.GetAccountHistoryInput{
		Range:        &r,
		Transactions: &d,
	}

	want := fmt.Sprintf("range=%s&transactions=%s", r, d)
	got := i.Encode()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestGetAccountHistory(t *testing.T) {
	t.Run("InvalidInput", func(t *testing.T) {
		a := allyinvest.NewWithClient(http.DefaultClient)
		_, err := a.GetAccountHistory(allyinvest.GetAccountHistoryInput{})
		if err == nil {
			t.Errorf("got nil, want err")
		}
	})

	t.Run("Success", func(t *testing.T) {
		a := allyinvest.NewWithClient(mockHTTPClient{
			DoValue: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte{})),
			},
		})

		_, err := a.GetAccountHistory(allyinvest.GetAccountHistoryInput{
			AccountID: allyinvest.String("s"),
		})

		if err == nil {
			t.Errorf("got nil, want err")
		}
	})
}

func TestGetAccount(t *testing.T) {
	t.Run("InvalidInput", func(t *testing.T) {
		a := allyinvest.NewWithClient(http.DefaultClient)
		_, err := a.GetAccount(allyinvest.GetAccountInput{})
		if err == nil {
			t.Errorf("got nil, want err")
		}
	})

	t.Run("Success", func(t *testing.T) {
		a := allyinvest.NewWithClient(mockHTTPClient{
			DoValue: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte{})),
			},
		})

		_, err := a.GetAccount(allyinvest.GetAccountInput{
			AccountID: allyinvest.String("s"),
		})

		if err == nil {
			t.Errorf("got nil, want err")
		}
	})
}

func TestGetAccountBalances(t *testing.T) {
	t.Run("InvalidInput", func(t *testing.T) {
		a := allyinvest.NewWithClient(http.DefaultClient)
		_, err := a.GetAccountBalances(allyinvest.GetAccountBalancesInput{})
		if err == nil {
			t.Errorf("got nil, want err")
		}
	})

	t.Run("Success", func(t *testing.T) {
		a := allyinvest.NewWithClient(mockHTTPClient{
			DoValue: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte{})),
			},
		})

		_, err := a.GetAccountBalances(allyinvest.GetAccountBalancesInput{
			AccountID: allyinvest.String("s"),
		})

		if err == nil {
			t.Errorf("got nil, want err")
		}
	})
}

func TestGetAccountsBalances(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		a := allyinvest.NewWithClient(mockHTTPClient{
			DoValue: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte{})),
			},
		})

		_, err := a.GetAccountsBalances()

		if err == nil {
			t.Errorf("got nil, want err")
		}
	})
}

func TestGetAccounts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		a := allyinvest.NewWithClient(mockHTTPClient{
			DoValue: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte{})),
			},
		})

		_, err := a.GetAccounts()

		if err == nil {
			t.Errorf("got nil, want err")
		}
	})
}
