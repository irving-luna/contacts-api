package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/unrolled/render"
)

var (
	crender = render.New()
)

// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func readBodyAs[T any](requestBody io.ReadCloser) (*T, error) {
	body := new(T)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(requestBody)
	// dec.DisallowUnknownFields()
	err := dec.Decode(&body)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			return nil, fmt.Errorf("request body contains badly-formed JSON (at position %d) %w", syntaxError.Offset, err)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return nil, fmt.Errorf("request body contains badly-formed JSON %w", err)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			return nil, fmt.Errorf("request body contains an invalid value for the %q field (at position %d) %w", unmarshalTypeError.Field, unmarshalTypeError.Offset, err)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return nil, fmt.Errorf("request body contains unknown field %s %w", fieldName, err)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			return nil, errors.New("request body must not be empty")

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			return nil, errors.New("request body must not be larger than 1MB")

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			return nil, fmt.Errorf("unhandle error %w", err)
		}
	}
	return body, nil
}