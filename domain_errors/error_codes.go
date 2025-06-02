package domain_errors

import "net/http"

var codes = map[int]error{}

var (
	ErrUnknown       = register(New(90000, http.StatusInternalServerError, "unknown error"))
	ErrDatabaseError = register(New(80000, http.StatusInternalServerError, "database error"))

	// General errors
	ErrNotFound           = register(New(10001, http.StatusNotFound, "not found"))
	ErrBadRequest         = register(New(10002, http.StatusBadRequest, "bad request"))
	ErrUnauthorized       = register(New(10003, http.StatusUnauthorized, "unauthorized"))
	ErrForbidden          = register(New(10004, http.StatusForbidden, "forbidden"))
	ErrTooManyRequests    = register(New(10005, http.StatusTooManyRequests, "too many requests"))
	ErrCSRFMismatch       = register(derive(10101, "CSRF token mismatch", ErrBadRequest))
	ErrBadArgument        = register(derive(10102, "bad argument", ErrBadRequest))
	ErrCantParseRequest   = register(derive(10103, "can not parse the request", ErrBadRequest))
	ErrRequestValidation  = register(derive(10104, "request validation error", ErrBadRequest))
	ErrSessionExpired     = register(derive(10201, "session expired", ErrUnauthorized))
	ErrUnableRefreshToken = register(derive(10202, "unable to refresh token", ErrUnauthorized))
	ErrTokenReuseDetected = register(derive(10203, "token reuse detected", ErrUnauthorized))

	ErrUserNotFound            = register(derive(20101, "user not found", ErrNotFound))
	ErrUserAlreadyExists       = register(derive(20102, "user already exists", ErrBadRequest))
	ErrEmailVerificationFailed = register(derive(20103, "wrong email verification code", ErrBadRequest))
	ErrUserBlocked             = register(derive(20104, "user is blocked", ErrForbidden))

	ErrAgentNotFound      = register(derive(20201, "agent not found", ErrNotFound))
	ErrAgentAlreadyExists = register(derive(20202, "agent already exists", ErrBadRequest))
	ErrAgentDisabled      = register(derive(20203, "agent is disabled", ErrForbidden))

	ErrProxyNotFound      = register(derive(20301, "proxy not found", ErrNotFound))
	ErrProxyAlreadyExists = register(derive(20302, "proxy already exists", ErrBadRequest))
	ErrNoProxiesAvailable = register(derive(20303, "no proxies available", ErrNotFound))
	ErrProxyDisabled      = register(derive(20304, "proxy is disabled", ErrForbidden))

	ErrHttpRouterNotFound      = register(derive(20401, "http router not found", ErrNotFound))
	ErrHttpRouterAlreadyExists = register(derive(20402, "http router already exists", ErrBadRequest))

	ErrTcpRouterNotFound      = register(derive(20501, "tcp router not found", ErrNotFound))
	ErrTcpRouterAlreadyExists = register(derive(20502, "tcp router already exists", ErrBadRequest))
)

func derive(code int, message string, baseErr error) error {
	err := MustUnwrapToDomainError(baseErr)
	return NewWrap(code, err.SuggestedHTTPStatus(), message, baseErr)
}

func register(err error) error {
	codes[MustUnwrapToDomainError(err).Code()] = err
	return err
}

func CodeToError(code int) error {
	if err, ok := codes[code]; ok {
		return err
	}
	return ErrUnknown
}
