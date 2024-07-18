package api

// Money Balance Params
type MoneyBalanceParams struct {
	Username string
}

// Money balance Response
type MoneyBalanceResponse struct {
	// Account Balance
	Balance int64

	// Error Message Code
	Code int
}

// Error Response
type ErrorResponse struct {
	// Error Message
	Message string

	// Error Message Code
	code int
}
