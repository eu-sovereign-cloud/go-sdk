package gosdk

import "context"

type TenantID string

type ctxTenantIDKeyType struct{}

var ctxTenantIDKey = ctxTenantIDKeyType{}

// WithTenantID returns a new context with the given TenantID.
func WithTenantID(ctx context.Context, tid TenantID) context.Context {
	return context.WithValue(ctx, ctxTenantIDKey, tid)
}

// TenantIDFromContext returns the TenantID from the given context.
func TenantIDFromContext(ctx context.Context) (TenantID, bool) {
	tid, ok := ctx.Value(ctxTenantIDKey).(TenantID)
	return tid, ok
}

// MustTenantIDFromContext returns the TenantID from the given context.
// If the TenantID is not found, it panics.
func MustTenantIDFromContext(ctx context.Context) TenantID {
	tid, ok := TenantIDFromContext(ctx)
	if !ok {
		panic(ErrNoTenantID)
	}
	return tid
}
