package api_models

import "github.com/pikvm/cloud-api/domain_errors"

// @description generic response model
type ResponseModel struct {
	Result  interface{}    `json:"result,omitempty" swaggertype:"object"`           // Response model. Can be null
	Message string         `json:"message,omitempty" example:"message from server"` // Optional message from server
	Error   *ResponseError `json:"error,omitempty"`                                 // Optional error
}

// @description response error model
type ResponseError struct {
	Code    int    `json:"code"`    // Error code
	Message string `json:"message"` // Error message
}

func (e *ResponseError) Error() string {
	return e.Message
}

func (e *ResponseError) ToDomainError() error {
	return domain_errors.CodeToError(e.Code)
}
