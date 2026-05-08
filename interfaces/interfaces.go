package interfaces

import (
	"io"
	"net/http"
)

type ClientInterface interface {
	Do(req *http.Request) (io.ReadCloser, error)
}
