package echoserver

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator"
)

type Server struct {
	server *echo.Echo
}

func NewServer(server *echo.Echo) server.Server {
	return &Server{server: server}
}

func (s *Server) Run(address string) {
	s.server.Logger.Fatal(s.server.Start(address))
}

func (s *Server) Debug(debug bool) {
	s.server.Debug = debug
}

func (s *Server) SetRequestValidator(validator validator.Validator) {
	s.server.Validator = validator
}
