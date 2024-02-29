package handler

import (
	"context"
	"fmt"
	"hexagonal/repository"
	"hexagonal/service"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PortfolioHandler struct {
		portServ service.PortfolioService
	}
)

// , portServ service.PortfolioService
func NewPortfolioHandler(app *echo.Echo) PortfolioHandler {

	// dbhost := "apicenter_mongo"
	dbhost := "localhost:27017"

	client, err := connectMongo(dbhost)
	if err != nil {
		panic(err)
	}

	portRepoMock := repository.NewPortfolioRepositoryMock()
	portRepoMongodb := repository.NewPortfolioRepositoryMongodb(client, "crypto")

	_ = portRepoMock
	_ = portRepoMongodb

	// fmt.Println("init with MongoDB")
	portServ := service.NewPortfolioService(portRepoMongodb)

	portRepoHandler := PortfolioHandler{portServ: portServ}

	app.GET("/portfolios", portRepoHandler.GetPortfolios)
	app.GET("/portfolios/:symbol", portRepoHandler.GetPortfolio)

	return portRepoHandler
}

func connectMongo(dbhost string) (client *mongo.Client, err error) {

	dbhost = fmt.Sprintf("mongodb://%s", dbhost)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client()
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(30)
	clientOptions.SetMaxConnIdleTime(0)
	clientOptions.SetMaxConnecting(3)
	clientOptions.ApplyURI(dbhost)
	tM := reflect.TypeOf(bson.M{})
	registry := bson.NewRegistryBuilder().RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM).Build()
	// registry := bson.NewRegistryBuilder().Build()
	clientOptions.SetRegistry(registry)
	client, err = mongo.Connect(ctx, clientOptions)

	return
}

func (h PortfolioHandler) GetPortfolios(ctx echo.Context) (err error) {
	userid := 1
	var ports service.PortfoliosResponse
	ports, err = h.portServ.GetPortfolios(userid)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"data": []string{}, "message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"data": ports})
}

func (h PortfolioHandler) GetPortfolio(ctx echo.Context) (err error) {
	symbol := ctx.Param("symbol")
	userid := 1
	var port *service.PortfolioResponse
	port, err = h.portServ.GetPortfolio(userid, symbol)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"data": map[string]string{}, "message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"data": port})
}
