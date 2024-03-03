package core

import "github.com/labstack/echo/v4"

var (
	_app *echo.Echo
)

func NewServer() *echo.Echo {
	if _app != nil {
		return _app
	}

	_app = echo.New()

	return _app
}
