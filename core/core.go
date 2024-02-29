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
	//app.GET("/", hello)

	//app.Logger.Fatal(app.Start(":8080"))
	return _app
}
