package translate

// ErrorModel used to response error
type ErrorModel struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Errors     interface{} `json:"errors,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
}
