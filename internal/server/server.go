package server

import (
	"fmt"
	"ticket/internal/auth"
	"ticket/internal/config"
	"ticket/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Server struct {
	db     database.Service
	router *echo.Echo
	config config.Config
}

func NewServer() *Server {
	cfg := config.New()
	db := database.New(cfg)
	server := &Server{
		db:     db,
		config: cfg,
		router: echo.New(),
	}
	auth.New(cfg)

	server.registerRoutes()
	server.configServer()

	return server
}

func (s *Server) Start() error {
	return s.router.Start(fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port))
}

func (s *Server) configServer() {
	s.router.HideBanner = true
	s.router.Logger.SetLevel(log.WARN)
}
