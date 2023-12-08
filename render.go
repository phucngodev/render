package render

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const (
	// request ID key in the request header.
	// normally this is provided by a request middleware.
	requestIdKey = "X-Request-Id"
	ctxKey       = "trace_id"
)

// getRequestId get trace_id from context, if no trace_id found
// get X-Request-Id from request header. otherwise generate new request ID
// trace_id, or X-Request-Id normally provides by tracing or requestId middleware.
func getRequestId(r *http.Request) string {
	requestId, ok := r.Context().Value(ctxKey).(string)
	if !ok || requestId == "" {
		requestId = r.Header.Get(requestIdKey)
	}

	if requestId == "" {
		requestId = uuid.NewString()
	}

	return requestId
}

// Success send json http success response.
func Success[T any](w http.ResponseWriter, r *http.Request, status int, v T) {
	requestId := getRequestId(r)
	resp := response[T]{
		Status:    statusSuccess,
		RequestID: requestId,
		Data:      v,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, err := Encode(resp)
	if err != nil {
		Error(w, r, http.StatusInternalServerError, "internal server error", err)
		return
	}

	w.WriteHeader(status)
	w.Write(body)
}

// Error send json http error response.
func Error(w http.ResponseWriter, r *http.Request, status int, msg string, err error) {
	requestId := getRequestId(r)
	resp := responseError{
		Status:    statusError,
		RequestID: requestId,
		Error:     msg,
	}

	log := logWithTrace(requestId)
	log.Error("response error", "error", fmt.Sprintf("%s - %v", msg, err))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	buf := &bytes.Buffer{}
	encoder := Encoder(buf)
	verr := encoder.Encode(resp)
	if verr != nil {
		log.Error("error encode response", "error", fmt.Sprintf("%v", err))
		w.WriteHeader(http.StatusInternalServerError)
		resp.Error = "internal server error"
		encoder.Encode(resp)
		w.Write(buf.Bytes())
		return
	}

	w.WriteHeader(status)
	w.Write(buf.Bytes())
}
