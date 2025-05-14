package routers

import (
	"github.com/iryoda/insecticide/src/controllers"
	"github.com/labstack/echo/v4"
)

type ComponentRouter struct {
	ComponentController controllers.ComponentController
}

func NewComponentRouter(sc controllers.ComponentController) ComponentRouter {
	return ComponentRouter{
		ComponentController: sc,
	}
}

func (sr *ComponentRouter) Register(e *echo.Echo) {
	g := e.Group("/components")

	g.POST("", sr.ComponentController.Create)
	g.POST("/:id", sr.ComponentController.GetById)
}
