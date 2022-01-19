package providers

import (
	"github.com/maurodanieldev/quasar-oper-fire/controllers"
	"github.com/maurodanieldev/quasar-oper-fire/router"
	"github.com/maurodanieldev/quasar-oper-fire/server"
	"github.com/maurodanieldev/quasar-oper-fire/services"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()
	_ = Container.Provide(server.NewServer)
	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(services.NewTrilaterationService)
	_ = Container.Provide(services.NewMessagesService)
	_ = Container.Provide(services.NewSatelliteService)

	_ = Container.Provide(controllers.NewTopSecretSplitHandler)
	_ = Container.Provide(controllers.NewTopSecretHandler)
	_ = Container.Provide(controllers.NewSatelliteHandler)
	return Container
}
