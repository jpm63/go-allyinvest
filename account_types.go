package allyinvest

import "time"

type GetAccountHoldingsOutput struct {
	Response GetAccountHoldingsResponse `json:"response"`
}

type GetAccountHoldingsResponse struct {
	ID              string          `json:"@id"`
	AccountHoldings AccountHoldings `json:"accountholdings"`
	ElapsedTime     int             `json:"elapsedtime,string"`
	Error           string          `json:"error"`
}

type AccountHoldings struct {
	Holding         []AccountHolding `json:"holding"`
	TotalSecurities float64          `json:"totalsecurities,string"`
}

type AccountHolding struct {
	AccountType       int        `json:"accounttype,string"`
	CostBasis         float64    `json:"costbasis,string"`
	GainLoss          float64    `json:"gainloss,string"`
	Instrument        Instrument `json:"instrument"`
	MarketValue       float64    `json:"marketvalue,string"`
	MarketValueChange float64    `json:"marketvaluechange,string"`
	Price             float64    `json:"price,string"`
	PurchasePrice     float64    `json:"purchaseprice,string"`
	Quantity          float64    `json:"qty,string"`
	Quote             Quote      `json:"quote"`
	SODCostBasis      int        `json:"sodcostbasis,string"`
	Underyling        string     `json:"underlying"`
}

const (
	AccountTypeCash        = 1
	AccountTypeMarginLong  = 2
	AccountTypeMarginShort = 5
)

type Instrument struct {
	CFI               string    `json:"cfi"`
	CUSIP             string    `json:"cusip"`
	Description       string    `json:"desc"`
	Factor            float64   `json:"factor,string"`
	MaturityDate      time.Time `json:"matdt"`
	MaturityMonthYear string    `json:"mmy"`
	Multiplier        int       `json:"mult,string"`
	PutCall           int       `json:"putcall,string"`
	SecurityType      string    `json:"sectyp"`
	StrikePrice       float64   `json:"strkpx,string"`
	Symbol            string    `json:"sym"`
}

const (
	InstrumentKindCommonStock = 0
	InstrumentKindOption      = 1
)

const (
	SecurityTypeOption      = "OPT"
	SecurityTypeCommonStock = "CS"
)

type Quote struct {
	Change        float64       `json:"change,string"`
	Format        string        `json:"format"`
	ExtendedQuote ExtendedQuote `json:"extendedquote"`
	LastPrice     float64       `json:"lastprice,string"`
}

type ExtendedQuote struct {
	DividentData string `json:"dividenddata"`
}

// AccountHistoryDateRange are the valid date range values.
type AccountHistoryDateRange string

const (
	AccountHistoryDateRangeAll          AccountHistoryDateRange = "all"
	AccountHistoryDateRangeToday        AccountHistoryDateRange = "today"
	AccountHistoryDateRangeCurrentWeek  AccountHistoryDateRange = "current_week"
	AccountHistoryDateRangeCurrentMonth AccountHistoryDateRange = "current_month"
	AccountHistoryDateRangeLastMonth    AccountHistoryDateRange = "last_month"
)

// AccountHistoryTransaction are the valid kinds of transactions.
type AccountHistoryTransaction string

const (
	AccountHistoryTransactionAll         AccountHistoryTransaction = "all"
	AccountHistoryTransactionBookkeeping AccountHistoryTransaction = "bookkeeping"
	AccountHistoryTransactionTrade       AccountHistoryTransaction = "trade"
)

type GetAccountHistoryOutput struct {
	Response GetAccountHistoryResponse `json:"response"`
}

type GetAccountHistoryResponse struct {
	ID           string       `json:"@id"`
	Transactions Transactions `json:"transactions"`
	ElapsedTime  int          `json:"elapsedtime,string"`
	Error        string       `json:"error"`
}

type Transactions struct {
	Transaction []Transaction `json:"transaction"`
}

type Transaction struct {
	Activity    string             `json:"activity"`
	Amount      float64            `json:"amount,string"`
	Date        time.Time          `json:"date"`
	Description string             `json:"desc"`
	Symbol      string             `json:"symbol"`
	Transaction TransactionDetails `json:"transaction"`
}

type TransactionDetails struct {
	AccountType     int       `json:"accounttype,string"`
	Commission      float64   `json:"commission,string"`
	Description     string    `json:"description"`
	Fee             float64   `json:"fee,string"`
	Price           float64   `json:"price,string"`
	Quantity        int       `json:"quantity,string"`
	SecurityFee     float64   `json:"secfee,string"`
	Security        Security  `json:"security"`
	SettlementDate  time.Time `json:"settlementdate"`
	Side            int       `json:"side,string"`
	Source          string    `json:"source"`
	TradeDate       time.Time `json:"tradedate"`
	TransactionID   int       `json:"transactionid,string"`
	TransactionType string    `json:"transactiontype"`
}

type Security struct {
	CUSIP        string `json:"cusip"`
	ID           string `json:"id"`
	SecurityType string `json:"sectyp"`
	Symbol       string `json:"sym"`
}

type GetAccountBalancesOutput struct {
	Response GetAccountBalancesResponse `json:"response"`
}

type GetAccountBalancesResponse struct {
	ID             string         `json:"@id"`
	AccountBalance AccountBalance `json:"accountbalance"`
	ElapsedTime    int            `json:"elapsedtime,string"`
	Error          string         `json:"error"`
}

type AccountBalance struct {
	Account      int         `json:"account,string"`
	AccountValue float64     `json:"accountvalue,string"`
	BuyingPower  BuyingPower `json:"buyingpower"`
	FedCall      string      `json:"fedcall"`
	HouseCall    string      `json:"housecall"`
	Money        Money       `json:"money"`
	Securities   Securities  `json:"securities"`
}

type BuyingPower struct {
	CashAvailableForWithdrawal float64 `json:"cashavailableforwithdrawal,string"`
	DayTrading                 float64 `json:"daytrading,string"`
	EquityPercentage           float64 `json:"equitypercentage,string"`
	Options                    float64 `json:"options,string"`
	SODDayTrading              float64 `json:"soddaytrading,string"`
	SODOptions                 float64 `json:"sodoptions,string"`
	SODStock                   float64 `json:"sodstock,string"`
	Stock                      float64 `json:"stock,string"`
}

type Money struct {
	AccruedInterest    float64 `json:"accruedinterest,string"`
	Cash               float64 `json:"cash,string"`
	CashAvailable      float64 `json:"cashavailable,string"`
	MarginBalance      float64 `json:"marginbalance,string"`
	MMF                float64 `json:"mmf,string"`
	Total              float64 `json:"total,string"`
	UnclearedPositions float64 `json:"uncleareddeposits,string"`
	UnsettledFunds     float64 `json:"unsettledfunds,string"`
	Yield              float64 `json:"yield,string"`
}

type Securities struct {
	LongOptions  float64 `json:"longoptions,string"`
	LongStocks   float64 `json:"longstocks,string"`
	Options      float64 `json:"options,string"`
	ShortOptions float64 `json:"shortoptions,string"`
	ShortStocks  float64 `json:"shortstocks,string"`
	Stocks       float64 `json:"stocks,string"`
	Total        float64 `json:"total,string"`
}

type GetAccountOutput struct {
	Response GetAccountResponse `json:"response"`
}

type GetAccountResponse struct {
	ID              string          `json:"@id"`
	AccountBalance  AccountBalance  `json:"accountbalance"`
	AccountHoldings AccountHoldings `json:"accountholdings"`
	ElapsedTime     int             `json:"elapsedtime,string"`
	Error           string          `json:"error"`
}

type GetAccountsBalancesOutput struct {
	Response GetAccountsBalancesResponse `json:"response"`
}

type GetAccountsBalancesResponse struct {
	ID             string                `json:"@id"`
	AccountBalance AccountBalanceSummary `json:"accountbalance"`
	TotalBalance   TotalBalance          `json:"totalbalance"`
	ElapsedTime    int                   `json:"elapsedtime,string"`
	Error          string                `json:"error"`
}

type AccountBalanceSummary struct {
	Account      int     `json:"account,string"`
	AccountName  string  `json:"accountname"`
	AccountValue float64 `json:"accountvalue,string"`
}

type TotalBalance struct {
	AccountValue float64 `json:"accountvalue,string"`
}

type GetAccountsOutput struct {
	Response GetAccountsResponse `json:"response"`
}

type GetAccountsResponse struct {
	ID          string   `json:"@id"`
	Accounts    Accounts `json:"accounts"`
	ElapsedTime int      `json:"elapsedtime,string"`
	Error       string   `json:"error"`
}

type AccountSummary struct {
	Account         int             `json:"account,string"`
	AccountBalance  AccountBalance  `json:"accountbalance"`
	AccountHoldings AccountHoldings `json:"accountholdings"`
}

type Accounts struct {
	AccountSummary AccountSummary `json:"accountsummary"`
}
