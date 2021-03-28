package allyinvest

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	validateTagKey   = "validate"
	validateTagValue = "required"
)

// validate ensures all required fields in provided struct
// are not nil.
func validate(i interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(i))
	t := reflect.TypeOf(i)

	for i := 0; i < t.NumField(); i++ {
		val, ok := t.Field(i).Tag.Lookup(validateTagKey)
		if !ok {
			continue
		}

		if strings.EqualFold(val, validateTagValue) && v.Field(i).IsZero() {
			return fmt.Errorf("field %s is required and cannot be nil", t.Field(i).Name)
		}
	}

	return nil
}

// String returns a pointer to s.
func String(s string) *string {
	return &s
}

// GoString returns the value held by s.
func GoString(s *string) string {
	return *s
}

// Int returns a pointer to i.
func Int(i int) *int {
	return &i
}

// GoInt returns the value held by i.
func GoInt(i *int) int {
	return *i
}

// Float64 returns a pointer to f.
func Float64(f float64) *float64 {
	return &f
}

// GoFloat64 returns the value held by f.
func GoFloat64(f *float64) float64 {
	return *f
}

// AccountHistoryDateRangePtr returns a pointer to a.
func AccountHistoryDateRangePtr(a AccountHistoryDateRange) *AccountHistoryDateRange {
	return &a
}

// GoAccountHistoryDateRange returns the value held by a.
func GoAccountHistoryDateRange(a *AccountHistoryDateRange) AccountHistoryDateRange {
	return *a
}

// AccountHistoryTransactionPtr returns a pointer to a.
func AccountHistoryTransactionPtr(a AccountHistoryTransaction) *AccountHistoryTransaction {
	return &a
}

// GoAccountHistoryTransaction returns the value held by a.
func GoAccountHistoryTransaction(a *AccountHistoryTransaction) AccountHistoryTransaction {
	return *a
}
