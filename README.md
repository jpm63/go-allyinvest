# Summary
Basic wrapper for the Ally Invest API. **WARNING - SOME OF THE APIS IN THIS MODULE LET YOU MAKE REAL TRADES WITH REAL MONEY. USE AT YOUR DISCRETION.**

# Usage
1. [Register](https://www.ally.com/api/invest/documentation/getting-started/) your OAUTH application with Ally and obtain the following keys:
    - Consumer Key
    - Consumer Secret
    - Access Token
    - Access Token Secret

2. Initialize the client and invoke endpoint:

```go
import "github.com/jpm63/go-allyinvest"

func main() {
    k := allyinvest.ApplicationKeys{
		ConsumerKey:    "CONSUMER_KEY",
		ConsumerSecret: "CONSUMER_SECRET",
		AccessToken:    "ACCESS_TOKEN",
		AccessSecret:   "ACCESS_SECRET",
	}

	api := allyinvest.New(k)

    // Invoke endpoint
    // out, err := api.SearchOptions(allyinvest.SearchOptionsInput{ ... })
}
```

# Completed APIS
- [X] Account
- [ ] Order/Trade
- [ ] Market (In Progress)
- [ ] Member
- [ ] Utility
- [ ] Watchlist
- [ ] Streaming