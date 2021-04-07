package allyinvest_test

import (
	"fmt"

	"github.com/jpm63/go-allyinvest"
)

func Example_aPIGetOptions() {
	k := allyinvest.ApplicationKeys{
		ConsumerKey:    "CONSUMER_KEY",
		ConsumerSecret: "CONSUMER_SECRET",
		AccessToken:    "ACCESS_TOKEN",
		AccessSecret:   "ACCESS_SECRET",
	}

	api := allyinvest.New(k)

	// Gets all SPY options expiring April 1, 2021 with a strike prices of 353.00
	out, err := api.SearchOptions(allyinvest.SearchOptionsInput{
		Symbol: allyinvest.String("SPY"),
		Query: &allyinvest.OptionSearchQuery{
			{
				ExpirationDate: allyinvest.Int(20210409),
				QueryOperator:  allyinvest.QueryOperatorEqual,
			},
			{
				StrikePrice:   allyinvest.Float64(353.00),
				QueryOperator: allyinvest.QueryOperatorEqual,
			},
		},
	})

	fmt.Println(err)
	fmt.Println(out)
}
