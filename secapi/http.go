package secapi

import (
	"net/http"
	"slices"
)

func checkStatusCode(code int, allowedCodes ...int) bool {
	return slices.Contains(allowedCodes, code)
}

func checkSuccessGetStatusCode(code int) bool {
	return checkStatusCode(code, http.StatusOK)
}

func checkSuccessPostStatusCode(code int) bool {
	return checkStatusCode(code, http.StatusAccepted)
}

func checkSuccessDeleteStatusCode(code int) bool {
	return checkStatusCode(code, http.StatusAccepted, http.StatusNotFound)
}

func checkSuccessPutStatusCode[T any](code int, createJSON, updateJSON *T) (bool, *T) {
	switch code {
	case http.StatusCreated:
		return true, createJSON
	case http.StatusOK:
		return true, updateJSON
	default:
		return false, nil
	}
}

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
	case http.StatusUnprocessableEntity:
		return ErrValidationError
	case http.StatusPreconditionFailed:
		return ErrRequestPreconditionFailed
	case http.StatusConflict:
		return ErrConflictingRequest
	case http.StatusInternalServerError:
		return ErrInternalError
	default:
		return ErrUnknownError
	}
}
