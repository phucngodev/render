package render

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/ajg/form"
)

// DecodeJSON decodes a given reader into an interface using the json decoder.
func DecodeJSON[T any](r io.Reader, v T) error {
	defer io.Copy(io.Discard, r)
	return json.NewDecoder(r).Decode(v)
}

// DecodeForm decodes a given reader into an interface using the form decoder.
func DecodeForm[T any](r io.Reader, v T) error {
	decoder := form.NewDecoder(r)
	return decoder.Decode(v)
}

// Encoder return json decoder.
// This should be used when your handler logic path require serveral encoding operations.
func Encoder(w io.Writer) *json.Encoder {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	return encoder
}

// Encode marshal input data v and return json byte array.
func Encode[T any](v T) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(true)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
