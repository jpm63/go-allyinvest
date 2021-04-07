package allyinvest

import (
	"fmt"
	"net/url"
	"strings"
)

// SearchOptionsInput are the required or optional parameters
// to SearchOptions.
type SearchOptionsInput struct {
	// REQUIRED: The symbol you are searching for.
	Symbol *string `validate:"required"`

	// OPTIONAL: Query to apply when searching for the provided symbol.
	Query *OptionSearchQuery

	// OPTIONAL: Fields to return from the search. If ommitted, all fields
	// are returned.
	Fields []string
}

// Encode encodes s to a query string.
func (s SearchOptionsInput) Encode() string {
	u := url.Values{}
	u.Add("symbol", *s.Symbol)

	if s.Query != nil {
		u.Add("query", s.Query.Encode())
	}

	if len(s.Fields) > 0 {
		u.Add("fids", strings.Join(s.Fields, ","))
	}

	return u.Encode()
}

// Encode encodes o into a custom query string.
func (o OptionSearchQuery) Encode() string {
	var t []string
	for _, e := range o {
		s := e.Encode()
		if s != "" {
			t = append(t, s)
		}
	}

	return strings.Join(t, " AND ")
}

// Encode encodes o into a custom query string.
func (o OptionSearchQueryElement) Encode() string {
	switch {
	case o.StrikePrice != nil:
		return fmt.Sprintf("strikeprice%s%.2f", o.QueryOperator, *o.StrikePrice)
	case o.ExpirationDate != nil:
		return fmt.Sprintf("xdate%s%d", o.QueryOperator, *o.ExpirationDate)
	case o.ExpirationMonth != nil:
		return fmt.Sprintf("xmonth%s%02d", o.QueryOperator, *o.ExpirationMonth)
	case o.ExpirationYear != nil:
		return fmt.Sprintf("xyear%s%d", o.QueryOperator, *o.ExpirationYear)
	case o.OptionKind != nil:
		return fmt.Sprintf("put_call%s%s", o.QueryOperator, *o.OptionKind)
	case o.Unique != nil:
		return fmt.Sprintf("unique%s%s", o.QueryOperator, *o.Unique)
	default:
		return ""
	}
}

// SearchOptions searches for the provided symbol.
func (a *api) SearchOptions(input SearchOptionsInput) (SearchOptionsOutput, error) {
	output := SearchOptionsOutput{}
	err := validate(input)
	if err != nil {
		return output, err
	}

	route := "/market/options/search"
	err = a.httpGet(route, input, &output)
	return output, err
}

// GetOptionsStrikesInput are the required or optional parameters
// to GetOptionsStrikes.
type GetOptionsStrikesInput struct {
	// REQUIRED: The symbol you are searching for.
	Symbol *string `validate:"required"`
}

// Encode encodes g to a query string.
func (g GetOptionsStrikesInput) Encode() string {
	u := url.Values{}
	if g.Symbol != nil {
		u.Add("symbol", string(*g.Symbol))
	}

	return u.Encode()
}

// GetOptionsStrikes lists the available strikes for the given symbol.
func (a *api) GetOptionsStrikes(input GetOptionsStrikesInput) (GetOptionsStrikesOutput, error) {
	output := GetOptionsStrikesOutput{}
	err := validate(input)
	if err != nil {
		return output, err
	}

	route := "/market/options/strikes"
	err = a.httpGet(route, input, &output)
	return output, err
}

// GetOptionsExpirationsInput are the required or optional parameters
// to GetOptiosnExpirations.
type GetOptionsExpirationsInput struct {
	// REQUIRED: The symbol you are searching for.
	Symbol *string `validate:"required"`
}

// Encode encodes g to a query string.
func (g GetOptionsExpirationsInput) Encode() string {
	u := url.Values{}
	if g.Symbol != nil {
		u.Add("symbol", string(*g.Symbol))
	}

	return u.Encode()
}

// GetOptionsStrikes lists the available strikes for the given symbol.
func (a *api) GetOptionsExpirations(input GetOptionsExpirationsInput) (GetOptionsExpirationsOutput, error) {
	output := GetOptionsExpirationsOutput{}
	err := validate(input)
	if err != nil {
		return output, err
	}

	route := "/market/options/expirations"
	err = a.httpGet(route, input, &output)
	return output, err
}
