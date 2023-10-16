package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
)

func Bind(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

type Responder struct {
	http.ResponseWriter
}

func (resp Responder) Respond(data any, err error) {
	status := http.StatusOK

	r := &Response{
		Data:   data,
		Errors: err,
	}

	if err != nil {
		status = http.StatusInternalServerError

		if er, ok := err.(*errors.Error); ok {
			status = er.StatusCode
		}
	}

	resp.ResponseWriter.Header().Set("Content-Type", "application/json")
	resp.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	resp.ResponseWriter.WriteHeader(status)

	json.NewEncoder(resp.ResponseWriter).Encode(r)
}

type Response struct {
	Data   any   `json:"data,omitempty"`
	Errors error `json:"errors,omitempty"`
}
