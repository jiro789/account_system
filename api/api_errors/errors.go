package api_errors

import "net/http"

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func InsufficientFundsError() *ApiError {
	return &ApiError{
		Status:  http.StatusBadRequest,
		Message: "insufficient funds to do the transaction",
		Code:    "insufficient_funds_error",
	}
}

func InvalidIdError() *ApiError {
	return &ApiError{
		Status:  http.StatusBadRequest,
		Message: "invalid id supplied",
		Code:    "invalid_id",
	}
}

func InvalidInputError() *ApiError {
	return &ApiError{
		Status:  http.StatusBadRequest,
		Message: "invalid input",
		Code:    "invalid_input",
	}
}

func TransactionNotFoundError() *ApiError {
	return &ApiError{
		Status:  http.StatusNotFound,
		Message: "The transaction was not found",
		Code:    "transaction_not_found",
	}
}
