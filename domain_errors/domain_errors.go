package domain_errors

import (
	"errors"
	"fmt"
	"net/http"
)

type DomainError interface {
	Error() string
	Code() int
	SuggestedHTTPStatus() int
	Unwrap() error
}

type errBase struct{}

func (e errBase) Error() string { return "Base error" }
func (e errBase) Is(target error) bool {
	_, ok1 := target.(errBase)
	_, ok2 := target.(*errBase)
	return ok1 || ok2
}

var ErrBase = &errBase{}

type domainErrorImpl struct {
	code                int
	suggestedHTTPStatus int
	message             string
	baseError           error
}

func (e domainErrorImpl) Error() string            { return e.message }
func (e domainErrorImpl) Code() int                { return e.code }
func (e domainErrorImpl) SuggestedHTTPStatus() int { return e.suggestedHTTPStatus }
func (e domainErrorImpl) Unwrap() error            { return e.baseError }
func (e domainErrorImpl) Is(target error) bool {
	if err, ok := target.(DomainError); ok {
		return e.code == err.Code()
	}
	return false
}
func New(code, suggestedHTTPStatus int, message string) error {
	return &domainErrorImpl{code, suggestedHTTPStatus, message, ErrBase}
}
func NewWrap(code, suggestedHTTPStatus int, message string, cause error) error {
	return &domainErrorImpl{code, suggestedHTTPStatus, message, cause}
}
func WrapCause(domainErr error, cause error) error {
	if cause == nil {
		return nil
	}
	unwrapped := MustUnwrapToDomainError(domainErr)
	return NewWrap(unwrapped.Code(), unwrapped.SuggestedHTTPStatus(), unwrapped.Error(), cause)
}
func WrapCauseWithMessage(domainError error, cause error) error {
	if cause == nil {
		return nil
	}
	unwrapped := MustUnwrapToDomainError(domainError)
	return NewWrap(
		unwrapped.Code(),
		unwrapped.SuggestedHTTPStatus(),
		fmt.Sprintf("%s: %s", unwrapped.Error(), cause.Error()),
		cause,
	)
}

func UnwrapToDomainError(err error) (domainErr DomainError, ok bool) {
	var ret *domainErrorImpl
	if errors.As(err, &ret) {
		return ret, true
	}
	return nil, false
}

func MustUnwrapToDomainError(err error) DomainError {
	if domainErr, ok := UnwrapToDomainError(err); ok {
		return domainErr
	}
	panic("domainErr can not be unwrapped to DomainError interface")
}

func SuggestHTTPStatus(err error) int {
	if unwrapped, ok := UnwrapToDomainError(err); ok {
		return unwrapped.SuggestedHTTPStatus()
	}
	return http.StatusInternalServerError
}

func IsDomainError(err error) bool {
	_, ok := UnwrapToDomainError(err)
	return ok
}
