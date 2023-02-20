package ginserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator"
)

type Server struct {
	server *gin.Engine
}

func NewServer(server *gin.Engine) server.Server {
	return &Server{server: server}
}

func (s *Server) Run(address string) {
	log.Fatal(s.server.Run(address))
}

func (s *Server) Debug(debug bool) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func (s *Server) SetRequestValidator(v validator.Validator) {
	if _, ok := binding.Validator.Engine().(validator.Validator); ok {
		log.Fatal("failed setting a custom validator for gin framework")
	}
}
