package allyinvest

import (
	"fmt"
	"net/url"
)

// GetAccountHodingsInput are the required or optional parameters
// to GetAccountHoldings.
type GetAccountHoldingsInput struct {
	// REQUIRED: The account ID for which you are requesting holdings.
	AccountID *string `validate:"required"`
}

// GetAccountHoldings returns the holdings for the specified account.
func (a *api) GetAccountHoldings(input GetAccountHoldingsInput) (GetAccountHoldingsOutput, error) {
	output := GetAccountHoldingsOutput{}
	err := validate(input)
	if err != nil {
		return output, err
	}

	route := fmt.Sprintf("/accounts/%s/holdings", *input.AccountID)
	err = a.httpGet(route, url.Values{}, &output)
	return output, err
}

// GetAccountHistoryInput are the required or optional parameters
// to GetAccountHistory.
type GetAccountHistoryInput struct {
	// REQUIRED: The account ID for which you are requesting history.
	AccountID *string `validate:"required"`

	// OPTIONAL: Date range of transacations. Valid values:
	//
	// AccountHistoryDateRangeAll: "all"
	// AccountHistoryDateRangeToday: "today"
	// AccountHistoryDateRangeCurrentWeek: "current_week"
	// AccountHistoryDateRangeCurrentMonth: "current_month"
	// AccountHistoryDateRangeLastMonth: "last_month"
	Range *AccountHistoryDateRange

	// OPTIONAL: The kind of transactions to show. Valid values:
	//
	// AccountHistoryTransactionAll: "all"
	// AccountHistoryTransactionBookkeeping: "bookkeeping"
	// AccountHistoryTransactionTrade: "trade"
	Transactions *AccountHistoryTransaction
}

// Encode encodes g to a query string.
func (g GetAccountHistoryInput) Encode() string {
	u := url.Values{}
	if g.Range != nil {
		u.Add("range", string(*g.Range))
	}

	if g.Transactions != nil {
		u.Add("transactions", string(*g.Transactions))
	}

	return u.Encode()
}

// GetAccountHistory returns the history for the specified account.
func (a *api) GetAccountHistory(input GetAccountHistoryInput) (GetAccountHistoryOutput, error) {
	output := GetAccountHistoryOutput{}
	err := validate(input)
	if err != nil {
		return output, err
	}

	route := fmt.Sprintf("/accounts/%s/history", *input.AccountID)
	err = a.httpGet(route, input, &output)
	return output, err
}

// GetAccountBalancesInput are the required or optional parameters
// to GetAccountBalances.
type GetAccountBalancesInput struct {
	// REQUIRED: The account ID for which you are requesting balances.
	AccountID *string `validate:"required"`
}

// GetAccountBalances returns the balances for the specified account.
func (a *api) GetAccountBalances(input GetAccountBalancesInput) (GetAccountBalancesOutput, error) {
	output := GetAccountBalancesOutput{}
	err := validate(input)
	if err != nil {
		return output, err
	}

	route := fmt.Sprintf("/accounts/%s/balances", *input.AccountID)
	err = a.httpGet(route, url.Values{}, &output)
	return output, err
}

// GetAccountInput are the required or optional parameters
// to GetAccount.
type GetAccountInput struct {
	// REQUIRED: The account ID for which you are requesting the summary.
	AccountID *string `validate:"required"`
}

// GetAccount returns the summary for the specified account.
func (a *api) GetAccount(input GetAccountInput) (GetAccountOutput, error) {
	output := GetAccountOutput{}
	err := validate(input)
	if err != nil {
		return output, err
	}

	route := fmt.Sprintf("/accounts/%s", *input.AccountID)
	err = a.httpGet(route, url.Values{}, &output)
	return output, err
}

// GetAccountsBalances returns the balance summary for each account.
func (a *api) GetAccountsBalances() (GetAccountsBalancesOutput, error) {
	output := GetAccountsBalancesOutput{}
	route := "/accounts/balances"
	err := a.httpGet(route, url.Values{}, &output)
	return output, err
}

// GetAccounts lists all accounts associated with a user.
func (a *api) GetAccounts() (GetAccountsOutput, error) {
	output := GetAccountsOutput{}
	route := "/accounts"
	err := a.httpGet(route, url.Values{}, &output)
	return output, err
}
