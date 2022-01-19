package interfaces

import "github.com/labstack/echo/v4"

type APM interface {
	Start()
	Stop()
	Middleware() echo.MiddlewareFunc
}