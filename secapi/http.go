package secapi

import (
	"net/http"
)

func mapStatusCodeToError(status int) error {
	switch status {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return ErrUnauthorizedAccess
	case http.StatusForbidden:
		return ErrForbiddenAccess
	case http.StatusNotFound:
		return ErrResourceNotFound
	case http.StatusBadRequest:
		return ErrInvalidRequest
	case http.StatusPreconditionFailed:
		return ErrRequestPreconditionFailed
	case http.StatusConflict:
		return ErrConflictingRequest
	case http.StatusInternalServerError:
		return ErrInternalError
	default:
		return ErrUnknowError
	}
}
