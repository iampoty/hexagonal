package main

import (
	"hexagonal/core"
	"hexagonal/handler"
)

func main() {

	// portRepoMock := repository.NewPortfolioRepositoryMock()
	// portRepoMongodb := repository.NewPortfolioRepositoryMongodb(client, "crypto")

	// // fmt.Println("init with MongoDB")
	// portServ := service.NewPortfolioService(portRepoMongodb)

	// fmt.Println("init with Mock")
	// portServ := service.NewPortfolioService(portRepoMock)

	app := core.NewServer()

	//portRepoHandler :=
	handler.NewPortfolioHandler(app)

	// _ = portRepoMock
	// _ = portRepoMongodb
	// _ app *echo.Echo= portRepoHandler

	//router.NewPortfolioRouter(app, portRepoHandler)

	app.Logger.Fatal(app.Start(":8080"))

}
