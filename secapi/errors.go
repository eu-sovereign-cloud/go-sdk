package secapi

import (
	"errors"
)

var (
	ErrProviderNotAvailable = errors.New("provider not available in the region")

	ErrNoMetadata          = errors.New("metadata is empty")
	ErrNoMetadataTenant    = errors.New("metadata tenant is empty")
	ErrNoMetadataWorkspace = errors.New("metadata workspace is empty")
	ErrNoMetadataNetwork   = errors.New("metadata network is empty")
	ErrNoMetadataName      = errors.New("metadata name is empty")

	ErrUnauthorizedAccess        = errors.New("unauthorized access")
	ErrForbiddenAccess           = errors.New("forbidden access")
	ErrResourceNotFound          = errors.New("resource not found")
	ErrInvalidRequest            = errors.New("invalid request")
	ErrRequestPreconditionFailed = errors.New("request precondition failed")
	ErrConflictingRequest        = errors.New("conflicting request")
	ErrInternalError             = errors.New("internal error")
	ErrUnknowError               = errors.New("unknow error")

	ErrRetryMaxAttemptsReached    = errors.New("max retry attempts reached")
	ErrRetryNotFoundExpectedValue = errors.New("not found the expected value")
	ErrRetryNotFoundExpectedError = errors.New("not found the expected error")
)
