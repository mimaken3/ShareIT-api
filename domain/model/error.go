package model

// errorResponse define struct hold data
// response error for client
type ErrorResponse struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}

// generalError interface should be implemented by errors that are to be handled by customers
type GeneralError interface {
	// Code return internal.ErrorCode to help customers figure out the abstract of the error
	Code() int
	// Messages returns error details to be shown to customers
	Messages() []string
}

// internalError interface should be implemented by errors that should be handled by service provider
// If there will be any necessity for categorization of internalErrors,
// i.e. automatic alert to different teams depending on error details
// `func Code() internal.ErrorCode` should be added to this interface at that time
type InternalError interface {
	// Implementation should simply be "return true"
	Internal() bool
}
