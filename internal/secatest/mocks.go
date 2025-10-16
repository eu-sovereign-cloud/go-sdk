package secatest

import (
	"encoding/json"
	"net/http"
)

const (
	headerContentTypeKey  = "Content-Type"
	headerContentTypeJSON = "application/json"
)

func encodeResponseBody(w http.ResponseWriter, data any) error {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

func configHttpResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set(headerContentTypeKey, headerContentTypeJSON)
	w.WriteHeader(statusCode)
}

func configGetHttpResponse(w http.ResponseWriter, data any) error {
	configHttpResponse(w, http.StatusOK)

	if err := encodeResponseBody(w, data); err != nil {
		return err
	}

	return nil
}

func configPutHttpResponse(w http.ResponseWriter, data any) error {
	configHttpResponse(w, http.StatusOK)

	if err := encodeResponseBody(w, data); err != nil {
		return err
	}

	return nil
}

func configPostHttpResponse(w http.ResponseWriter) {
	configHttpResponse(w, http.StatusAccepted)
}

func configDeleteHttpResponse(w http.ResponseWriter) {
	configHttpResponse(w, http.StatusAccepted)
}
