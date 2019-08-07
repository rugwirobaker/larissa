package handlers_test

import (
	"testing"

	"github.com/rugwirobaker/larissa/pkg/handlers"
	"github.com/rugwirobaker/larissa/pkg/larissa"
)

func newService() larissa.Service {
	return nil
}

func newHTTPHandler(svc larissa.Service) handlers.HTTPHandler {
	return handlers.NewHTTPHandler(svc)
}

func TestPutHandler(t *testing.T) {}

func TestGetHandler(t *testing.T) {}

func TestDelHandler(t *testing.T) {}

func TestExistsHandler(t *testing.T) {}

func TestErrorHandler(t *testing.T) {}
