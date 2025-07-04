package secatest

import (
	"bytes"
	"net/http"
	"text/template"
)

func processTemplateResponse(w http.ResponseWriter, name string, data any) {
	tmpl := template.Must(template.New("response").Parse(name))

	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, data)
	_, _ = w.Write(buffer.Bytes())
}

func configMockResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	w.WriteHeader(statusCode)
}

func configTemplateMockResponse(w http.ResponseWriter, statusCode int, template string, data any) {
	configMockResponse(w, statusCode)
	processTemplateResponse(w, template, data)
}
