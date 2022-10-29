package gh

import (
	"io"

	"github.com/cli/go-gh"
)

//go:generate moq -rm -out rest_client_mock.go . RESTClient

type RESTClient interface {
	Delete(path string, response interface{}) error
	Get(path string, response interface{}) error
	Patch(path string, body io.Reader, response interface{}) error
	Post(path string, body io.Reader, response interface{}) error
	Put(path string, body io.Reader, response interface{}) error
}

func NewRESTClient() (RESTClient, error) { //nolint: ireturn
	return gh.RESTClient(nil)
}
