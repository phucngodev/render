package render

const (
	statusSuccess = "success"
	statusError   = "error"
)

// M is a convenience alias for quickly building map type
// for response.
type M map[string]any

// response defines success response object structure.
type response[T any] struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id,omitempty"`
	Data      T      `json:"data,omitempty"`
}

// responseError defines error response object structure.
type responseError struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id,omitempty"`
	Error     string `json:"error"`
}
