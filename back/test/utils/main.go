package test_utils

import (
	"net"

	"github.com/labstack/echo/v4"
)

func SetRandomPortEcho(e *echo.Echo) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Listener = listener
}
