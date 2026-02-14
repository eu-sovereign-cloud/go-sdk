package secapi

import "errors"

var (
	ErrNoMetatada          = errors.New("metadata is empty")
	ErrNoMetatadaTenant    = errors.New("metadata tenant is empty")
	ErrNoMetatadaWorkspace = errors.New("metadata workspace is empty")
	ErrNoMetatadaNetwork   = errors.New("metadata network is empty")
	ErrNoMetatadaName      = errors.New("metadata name is empty")

	ErrUnauthorizedAccess        = errors.New("unauthorized access")
	ErrForbiddenAccess           = errors.New("forbidden access")
	ErrResourceNotFound          = errors.New("resource not found")
	ErrInvalidRequest            = errors.New("invalid request")
	ErrRequestPreconditionFailed = errors.New("request precondition failed")
	ErrConflictingRequest        = errors.New("conflicting request")
	ErrInternalError             = errors.New("internal error")
	ErrUnknowError               = errors.New("unknow error")

	ErrRetryMaxAttemptsReached = errors.New("max retry attempts reached")
)
