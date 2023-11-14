package handler

type Response struct {
	Data   any   `json:"data,omitempty"`
	Errors error `json:"errors,omitempty"`
}
