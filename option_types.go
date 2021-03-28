package allyinvest

import "time"

type OptionSearchQuery []OptionSearchQueryElement

type OptionSearchQueryElement struct {
	StrikePrice     *float64
	ExpirationDate  time.Time
	ExpirationMonth time.Time
	ExpirationYear  time.Time
	OptionKind      *string
	Unique          *string
	QueryOperator   QueryOperator
}

const (
	OptionKindPut  = "put"
	OptionKindCall = "call"
)

const (
	UniqueStrikePrice    = "strikeprice"
	UniqueExpirationDate = "xdate"
)

type QueryOperator string

const (
	QueryOperatorLessThan           QueryOperator = "-lt:"
	QueryOperatorGreaterThan        QueryOperator = "-gt:"
	QueryOperatorLessThanOrEqual    QueryOperator = "-lte:"
	QueryOperatorGreaterThanOrEqual QueryOperator = "-gte:"
	QueryOperatorEqual              QueryOperator = "-eq:"
)

type SearchOptionsOutput struct {
	Response SearchOptionsResponse `json:"response"`
}

type SearchOptionsResponse struct {
	ID          string        `json:"@id"`
	Quotes      OptionsQuotes `json:"quotes"`
	ElapsedTIme int           `json:"elapsedtime,string"`
	Error       string        `json:"error"`
}

type OptionsQuotes struct {
	Quote []OptionQuote `json:"quote"`
}

type OptionQuote struct {
	Ask                        float64   `json:"ask,string"`
	AskTime                    string    `json:"ask_time"`
	AskSize                    int       `json:"asksz,string"`
	Basis                      string    `json:"basis"`
	Bid                        float64   `json:"bid,string"`
	BidTime                    string    `json:"bid_time"`
	BidSize                    int       `json:"bidsz,string"`
	Change                     float64   `json:"chg,string"`
	ChangeDirection            string    `json:"chg_sign"`
	ChangeText                 string    `json:"chg_t"`
	Close                      float64   `json:"cl,string"`
	ContractSize               int       `json:"contract_size,string"`
	Date                       string    `json:"date"`
	DateTime                   time.Time `json:"datetime"`
	DaysToExpiration           int       `json:"days_to_expiration,string"`
	Exchange                   string    `json:"exch"`
	ExchangeDescription        string    `json:"exch_desc"`
	DayHigh                    float64   `json:"hi,string"`
	Delta                      float64   `json:"idelta,string"`
	Gamma                      float64   `json:"igamma,string"`
	ImpliedVolatility          float64   `json:"imp_Volatility,string"`
	LastTradeVolume            int       `json:"incr_vl,string"`
	Rho                        float64   `json:"irho,string"`
	Description                string    `json:"issue_desc"`
	Theta                      float64   `json:"itheta,string"`
	Vega                       float64   `json:"ivega,string"`
	Last                       float64   `json:"last,string"`
	DayLow                     float64   `json:"lo,string"`
	SettlementDesignation      string    `json:"op_delivery"`
	HasOptions                 int       `json:"op_flag,string"`
	OptionStyle                string    `json:"op_style"`
	OptionSubclass             int       `json:"ob_subclass,string"`
	OpenInterest               int       `json:"openinterest,string"`
	Open                       float64   `json:"opn,string"`
	PercentChange              float64   `json:"pchg,string"`
	PercentChangeSign          string    `json:"pchg_sign"`
	PriorDayClose              float64   `json:"pcls,string"`
	PriorDayHigh               float64   `json:"phi,string"`
	PriorDayLow                float64   `json:"plo,string"`
	PriorDayOpen               float64   `json:"popn,string"`
	PriorLastDate              string    `json:"pr_date"`
	PriorOpenInterest          string    `json:"pr_openinterest"`
	PriorChange                float64   `json:"prchg,string"`
	PremiumMultiplier          int       `json:"prem_mult,string"`
	PriorDatyVolume            int       `json:"pvol,string"`
	RootSymbol                 string    `json:"rootsymbol"`
	SecurityClass              int       `json:"secclass,string"`
	Session                    string    `json:"sesn"`
	StrikePrice                float64   `json:"strikeprice,string"`
	Symbol                     string    `json:"symbol"`
	TradeCondition             string    `json:"tcond"`
	Timestamp                  int       `json:"timestamp,string"`
	NumberTradesSinceOpen      int       `json:"tr_num,string"`
	Tradetick                  string    `json:"tradetick"`
	UnderlyingCUSIP            string    `json:"under_cusip"`
	UnderlyingSymbol           string    `json:"undersymbol"`
	Volume                     int       `json:"vl,string"`
	VolumeWeightedAveragePrice string    `json:"vwap"`
	Week52High                 float64   `json:"wk52hi,string"`
	Week52HighDate             string    `json:"wk52hidate"`
	Week52Low                  float64   `json:"wk52lo,string"`
	Week52LowDate              string    `json:"wk52lodate"`
	ExpirationDate             string    `json:"xdate"`
	ExpirationDay              string    `json:"xday"`
	ExpirationMonth            string    `json:"xmonth"`
	Expirationyear             string    `json:"xyear"`
}

const (
	ChangeDirectionEven = "e"
	ChangeDirectionUp   = "u"
	ChangeDirectionDown = "d"
)

const (
	SettlementDesignationStandard    = "S"
	SettlementDesignationNonStandard = "N"
	SettlementDesignationNA          = "X"
)

const (
	OptionStyleAmerican = "A"
	OptionStyleEuropean = "E"
)

const (
	OptionSubclassStandard  = 0
	OptionSubclassLeap      = 1
	OptionSubclassShortTerm = 3
)

const (
	SecurityClassStock  = 0
	SecurityClassOption = 1
)

const (
	SessionPreMarket  = "pre"
	SessionOpenMarket = "open"
	SessionPostMarket = "post"
)
