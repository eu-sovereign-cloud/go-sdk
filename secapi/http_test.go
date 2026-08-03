package secapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mapStatusCodeToError must map every status the ecp gateway actually returns to a distinct,
// named error. 422 Unprocessable Entity is returned for domain validation failures (bad CIDR,
// missing required fields, malformed references, ...) — a routine failure mode, not an unknown
// one — and previously fell through to the generic ErrUnknownError, discarding the fact that the
// request was rejected as invalid.
func TestMapStatusCodeToError(t *testing.T) {
	testCases := []struct {
		name   string
		status int
		want   error
	}{
		{"OK maps to nil", http.StatusOK, nil},
		{"Unauthorized", http.StatusUnauthorized, ErrUnauthorizedAccess},
		{"Forbidden", http.StatusForbidden, ErrForbiddenAccess},
		{"NotFound", http.StatusNotFound, ErrResourceNotFound},
		{"BadRequest", http.StatusBadRequest, ErrInvalidRequest},
		{"UnprocessableEntity maps to ErrValidationError, not ErrUnknowError", http.StatusUnprocessableEntity, ErrValidationError},
		{"PreconditionFailed", http.StatusPreconditionFailed, ErrRequestPreconditionFailed},
		{"Conflict", http.StatusConflict, ErrConflictingRequest},
		{"InternalServerError", http.StatusInternalServerError, ErrInternalError},
		{"unmapped status falls back to ErrUnknowError", http.StatusTeapot, ErrUnknownError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, mapStatusCodeToError(tc.status))
		})
	}
}
