package secatest

import (
	"bytes"
	"net/http"
	"text/template"

	mockRegion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"

	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/stretchr/testify/mock"
)

// Region
func MockListRegionsV1(sim *mockRegion.MockServerInterface, resp ListRegionsResponseV1) {
	json := template.Must(template.New("response").Parse(ListRegionsResponseTemplateV1))

	sim.EXPECT().ListRegions(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, lrp region.ListRegionsParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		for i := range resp.Providers {
			resp.Providers[i].URL = "http://" + r.Host + resp.Providers[i].URL
		}

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetRegionV1(sim *mockRegion.MockServerInterface, resp GetRegionResponseV1) {
	json := template.Must(template.New("response").Parse(GetRegionResponseTemplateV1))

	sim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, name string) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		for i := range resp.Providers {
			resp.Providers[i].URL = "http://" + r.Host + resp.Providers[i].URL
		}

		writeTemplateResponse(w, json, resp)
	})
}

func writeTemplateResponse(w http.ResponseWriter, tmpl *template.Template, data any) {
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, data)
	_, _ = w.Write(buffer.Bytes())
}
