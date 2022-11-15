package server_test

import (
	"github.com/halilylm/ticketing-ticket/test/integration"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestServer_Run(t *testing.T) {
	integration.CreateServer()
	res, err := http.Get("http://localhost:5000")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
