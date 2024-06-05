package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func (s *Server) AuthRedirect(c echo.Context) error {
	res := c.Response().Writer
	req := gothic.GetContextWithProvider(c.Request(), "openid-connect")
	if user, err := gothic.CompleteUserAuth(res, req); err == nil {
		return c.JSON(http.StatusOK, user)
	}
	gothic.BeginAuthHandler(res, req)
	return nil
}

func (s *Server) AuthCallback(c echo.Context) error {
	res := c.Response().Writer
	req := c.Request()
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}
