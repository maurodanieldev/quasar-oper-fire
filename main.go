package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/maurodanieldev/quasar-oper-fire/apm"
	"github.com/maurodanieldev/quasar-oper-fire/config"
	_ "github.com/maurodanieldev/quasar-oper-fire/docs"
	"github.com/maurodanieldev/quasar-oper-fire/providers"
	"github.com/maurodanieldev/quasar-oper-fire/router"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

// @title Quasar Fire API
// @version 1.0
// @description Operación Fuego de Quasar.
// @contact.name Daniel Mauro
// @contact.email mauro.daniel.dev@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/quasar-fire
func main() {

	apm.Get().Start()
	defer apm.Get().Stop()
	container := providers.BuildContainer()

	err := container.Invoke(func(server *echo.Echo, route *router.Router) {
		server.Debug = config.Environments().Postfix == "dev"
		route.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		panic(err)
	}

}
