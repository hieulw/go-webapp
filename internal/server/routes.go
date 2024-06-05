package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) registerRoutes() {
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())

	s.router.GET("/", s.HelloWorldHandler)
	s.router.GET("/health", s.healthHandler)

	s.router.GET("/auth/:provider", s.AuthRedirect)
	s.router.GET("/auth/callback", s.AuthCallback)
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
