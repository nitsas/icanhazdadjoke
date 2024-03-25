package client

import (
	"fmt"
	"net/http"
)

type ResponseStatusError struct {
	StatusCode int
}

func (e *ResponseStatusError) Error() string {
	return fmt.Sprintf("Got response status %d (%s)", e.StatusCode, http.StatusText(e.StatusCode))
}
