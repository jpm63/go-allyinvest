package allyinvest_test

import (
	"testing"

	"github.com/jpm63/go-allyinvest"
)

func TestNew(t *testing.T) {
	allyinvest.New(allyinvest.ApplicationKeys{})
}
