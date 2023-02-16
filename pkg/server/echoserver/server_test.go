package echoserver

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewServer_whenServerIsNotNilAndSameType(t *testing.T) {
	echoServer := NewServer(echo.New())

	assert.NotNil(t, echoServer)
}
