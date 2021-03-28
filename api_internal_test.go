package allyinvest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type mockHTTPClient struct {
	DoValue *http.Response
	DoError error
}

func (m mockHTTPClient) Do(*http.Request) (*http.Response, error) {
	return m.DoValue, m.DoError
}

type mockReadCloser struct {
	ReadValue  int
	ReadError  error
	CloseError error
}

func (m mockReadCloser) Read(p []byte) (int, error) {
	return m.ReadValue, m.ReadError
}

func (m mockReadCloser) Close() error {
	return m.CloseError
}

func TestAPIHTTPGet(t *testing.T) {
	t.Run("DoError", func(t *testing.T) {
		m := mockHTTPClient{
			DoError: fmt.Errorf("error"),
		}

		a := &api{
			client: m,
		}

		got := a.httpGet("", url.Values{}, nil)
		if got == nil {
			t.Errorf("got nil, want err")
		}
	})

	t.Run("ReadBodyError", func(t *testing.T) {
		m := mockReadCloser{
			ReadError: fmt.Errorf("error"),
		}

		n := mockHTTPClient{
			DoValue: &http.Response{
				Body: m,
			},
		}

		a := &api{
			client: n,
		}

		got := a.httpGet("", url.Values{}, nil)
		if got == nil {
			t.Errorf("got nil, want err")
		}
	})

	t.Run("Success", func(t *testing.T) {
		m := mockHTTPClient{
			DoValue: &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"key": 0}`))),
			},
		}

		a := &api{
			client: m,
		}

		var o map[string]interface{}
		got := a.httpGet("", url.Values{
			"key": []string{"value"},
		}, &o)
		if got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})
}

/*
func TestGetAccounts(t *testing.T) {
	k := allyinvest.ApplicationKeys{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		AccessToken:    os.Getenv("ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("ACCESS_SECRET"),
	}

	api, err := allyinvest.New(k)
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	err = api.GetAccounts()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
}
*/
