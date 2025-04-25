package xcoin_client

import "fmt"

// HTTP Status Codes
const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500
)

// Common errors
var (
	// ErrInvalidResponse indicates that the server response could not be parsed
	ErrInvalidResponse = fmt.Errorf("invalid response from server")

	// ErrUnexpectedStatusCode indicates that the server returned an unexpected status code
	ErrUnexpectedStatusCode = fmt.Errorf("unexpected status code from server")

	// ErrRequestCreation indicates that the HTTP request could not be created
	ErrRequestCreation = fmt.Errorf("failed to create HTTP request")

	// ErrRequestExecution indicates that the HTTP request could not be executed
	ErrRequestExecution = fmt.Errorf("failed to execute HTTP request")

	// ErrHTMLParsing indicates that the HTML response could not be parsed
	ErrHTMLParsing = fmt.Errorf("failed to parse HTML response")

	// ErrJSONDecoding indicates that the JSON response could not be decoded
	ErrJSONDecoding = fmt.Errorf("failed to decode JSON response")

	// ErrJSONEncoding indicates that the JSON request could not be encoded
	ErrJSONEncoding = fmt.Errorf("failed to encode JSON request")
)
