package allyinvest_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/jpm63/go-allyinvest"
)

func TestSearchOptionsInputEncode(t *testing.T) {
	s := allyinvest.SearchOptionsInput{
		Symbol: allyinvest.String("SPY"),
		Query: &allyinvest.OptionSearchQuery{
			{
				StrikePrice:   allyinvest.Float64(1),
				QueryOperator: allyinvest.QueryOperatorEqual,
			},
		},
		Fields: []string{"one", "two"},
	}

	got, err := url.ParseQuery(s.Encode())
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	sym := got["symbol"][0]
	if sym != "SPY" {
		t.Errorf("got %s, want SPY", sym)
	}

	query := got["query"][0]
	if query != "strikeprice-eq:1.00" {
		t.Errorf("got %s, want strikeprice-eq:1.00", query)
	}

	fields := got["fids"][0]
	if fields != "one,two" {
		t.Errorf("got %s, want one,two", fields)
	}
}

func TestOptionSearchQueryEncode(t *testing.T) {
	o := allyinvest.OptionSearchQuery{
		allyinvest.OptionSearchQueryElement{
			StrikePrice:   allyinvest.Float64(1),
			QueryOperator: allyinvest.QueryOperatorEqual,
		},
		allyinvest.OptionSearchQueryElement{},
		allyinvest.OptionSearchQueryElement{
			StrikePrice:   allyinvest.Float64(5.50),
			QueryOperator: allyinvest.QueryOperatorEqual,
		},
	}

	want := "strikeprice-eq:1.00 AND strikeprice-eq:5.50"
	got := o.Encode()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestOptionSearchQueryElementEncode(t *testing.T) {
	d, _ := time.Parse("2006-Jan-02", "2010-Jul-01")
	fmt.Println(d)
	tests := []struct {
		element allyinvest.OptionSearchQueryElement
		want    string
	}{
		{
			element: allyinvest.OptionSearchQueryElement{
				StrikePrice:   allyinvest.Float64(5),
				QueryOperator: allyinvest.QueryOperatorEqual,
			},
			want: "strikeprice-eq:5.00",
		},
		{
			element: allyinvest.OptionSearchQueryElement{
				ExpirationDate: d,
				QueryOperator:  allyinvest.QueryOperatorGreaterThan,
			},
			want: "xdate-gt:20100701",
		},
		{
			element: allyinvest.OptionSearchQueryElement{
				ExpirationMonth: d,
				QueryOperator:   allyinvest.QueryOperatorGreaterThanOrEqual,
			},
			want: "xmonth-gte:07",
		},
		{
			element: allyinvest.OptionSearchQueryElement{
				ExpirationYear: d,
				QueryOperator:  allyinvest.QueryOperatorLessThan,
			},
			want: "xyear-lt:2010",
		},
		{
			element: allyinvest.OptionSearchQueryElement{
				OptionKind:    allyinvest.String("put"),
				QueryOperator: allyinvest.QueryOperatorLessThanOrEqual,
			},
			want: "put_call-lte:put",
		},
		{
			element: allyinvest.OptionSearchQueryElement{
				Unique:        allyinvest.String("xdate"),
				QueryOperator: allyinvest.QueryOperatorEqual,
			},
			want: "unique-eq:xdate",
		},
		{
			element: allyinvest.OptionSearchQueryElement{},
			want:    "",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			got := tt.element.Encode()
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSearchOptions(t *testing.T) {
	t.Run("InvalidInput", func(t *testing.T) {
		a := allyinvest.NewWithClient(http.DefaultClient)
		_, err := a.SearchOptions(allyinvest.SearchOptionsInput{})
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

		_, err := a.SearchOptions(allyinvest.SearchOptionsInput{
			Symbol: allyinvest.String("s"),
		})

		if err == nil {
			t.Errorf("got nil, want err")
		}
	})
}
