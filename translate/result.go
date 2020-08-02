package translate

// ResultSuccess used to response error
type ResultSuccess struct {
	Status     string      `json:"status" example:"OK"`
	StatusCode int         `json:"status_code" example:"200"`
	Data       interface{} `json:"data,omitempty"`
}

// ResultErrors show you error out
type ResultErrors struct {
	Status     string      `json:"status" example:"Bad Request"`
	StatusCode int         `json:"status_code" example:"400"`
	Errors     interface{} `json:"errors,omitempty" swaggerignore:"true"`
	Message    string      `json:"message,omitempty" example:"Dynamic message"`
}
