package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maurodanieldev/quasar-oper-fire/apm"
	"github.com/maurodanieldev/quasar-oper-fire/controllers"
	_ "github.com/maurodanieldev/quasar-oper-fire/docs"
	"github.com/maurodanieldev/quasar-oper-fire/enums"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server                *echo.Echo
	topSecretHandler      *controllers.TopSecretHandler
	topSecretSplitHandler *controllers.TopSecretSplitHandler
	satelliteHandler      *controllers.SatelliteHandler
}

func NewRouter(
	server *echo.Echo,
	messagesHandler *controllers.TopSecretHandler,
	coordinatesHandler *controllers.TopSecretSplitHandler,
	satelliteHandler *controllers.SatelliteHandler,
) *Router {
	return &Router{
		server:                server,
		topSecretHandler:      messagesHandler,
		topSecretSplitHandler: coordinatesHandler,
		satelliteHandler:      satelliteHandler,
	}
}

func (r *Router) Init() {
	r.server.Pre(middleware.RemoveTrailingSlash())
	r.server.Use(middleware.Recover())
	r.server.Use(middleware.Logger())
	if apmMiddleware := apm.Get().Middleware(); apmMiddleware != nil {
		r.server.Use(apmMiddleware)
	}
	apiGroup := r.server.Group(enums.BasePath)
	{
		apiGroup.GET(enums.HealthRoute, controllers.HealthCheck)
		apiGroup.GET(enums.Swagger, echoSwagger.WrapHandler)
		messageGroup := apiGroup.Group(enums.TopSecret)
		{
			messageGroup.POST(enums.EmptyRoute, r.topSecretHandler.GetMessages)
		}
		coordinatesGroup := apiGroup.Group(enums.TopSecretSplit)
		{
			coordinatesGroup.GET(enums.EmptyRoute, r.topSecretSplitHandler.GetMessage)
			topSecretName := coordinatesGroup.Group("/:name")
			topSecretName.POST(enums.EmptyRoute, r.topSecretSplitHandler.PutMessageOnSatellite)
		}
		satellites := apiGroup.Group(enums.Satellites)
		{
			satellites.GET(enums.EmptyRoute, r.satelliteHandler.SatellitesGetAll)
			satellites.POST(enums.EmptyRoute, r.satelliteHandler.SatellitesPostOne)

			satellitesName := satellites.Group("/:name")
			satellitesName.GET(enums.EmptyRoute, r.satelliteHandler.SatellitesGetOne)
			satellitesName.PUT(enums.EmptyRoute, r.satelliteHandler.SatellitesPutOne)
			satellitesName.PATCH(enums.EmptyRoute, r.satelliteHandler.SatellitesPatchOne)
			satellitesName.DELETE(enums.EmptyRoute, r.satelliteHandler.SatellitesDeleteOne)
		}

	}
}
