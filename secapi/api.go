package secapi

import (
	"context"
	"net/http"
)

type API struct {
	authToken string
}

func (api *API) loadRequestHeaders(ctx context.Context, req *http.Request) error {
	req.Header.Set(headerKeyAuthorization, api.authToken)
	req.Header.Set(headerKeyAccept, headerValueAcceptJSON)
	return nil
}
