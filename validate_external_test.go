package allyinvest_test

import (
	"testing"

	"github.com/jpm63/go-allyinvest"
)

func TestString(t *testing.T) {
	want := "s"
	got := allyinvest.String(want)

	if *got != want {
		t.Errorf("got %s, want %s", *got, want)
	}
}

func TestGoString(t *testing.T) {
	want := "s"
	got := allyinvest.GoString(&want)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestInt(t *testing.T) {
	want := 1
	got := allyinvest.Int(want)

	if *got != want {
		t.Errorf("got %d, want %d", *got, want)
	}
}

func TestGoInt(t *testing.T) {
	want := 1
	got := allyinvest.GoInt(&want)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFloat64(t *testing.T) {
	want := float64(1)
	got := allyinvest.Float64(want)

	if *got != want {
		t.Errorf("got %f, want %f", *got, want)
	}
}

func TestGoFloat64(t *testing.T) {
	want := float64(1)
	got := allyinvest.GoFloat64(&want)

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestAccountHistoryDateRangePtr(t *testing.T) {
	want := allyinvest.AccountHistoryDateRangeAll
	got := allyinvest.AccountHistoryDateRangePtr(want)

	if *got != want {
		t.Errorf("got %s, want %s", *got, want)
	}
}

func TestGoAccountHistoryDateRange(t *testing.T) {
	want := allyinvest.AccountHistoryDateRangeAll
	got := allyinvest.GoAccountHistoryDateRange(&want)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAccountHistoryTransactionPtr(t *testing.T) {
	want := allyinvest.AccountHistoryTransactionAll
	got := allyinvest.AccountHistoryTransactionPtr(want)

	if *got != want {
		t.Errorf("got %s, want %s", *got, want)
	}
}

func TestGoAccountHistoryTransaction(t *testing.T) {
	want := allyinvest.AccountHistoryTransactionAll
	got := allyinvest.GoAccountHistoryTransaction(&want)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
