package main

import (
	"fmt"
	"hexagonal/core"
	"hexagonal/handler"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var (
	APP_MOCK = os.Getenv("APP_MOCK") == "1"
	app      *echo.Echo
)

func main() {

	log.Println("APP_MOCK: ", APP_MOCK)

	initTimezone()
	initConfig()

	app = core.NewServer()

	app.GET("/app/_router", routeInfoHandler)

	handler.NewPortfolioHandler(app)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", viper.GetInt("app.port"))))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func routeInfoHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, app.Routes())
}
