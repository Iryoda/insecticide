package routers

import (
	"github.com/iryoda/insecticide/src/controllers"
	"github.com/labstack/echo/v4"
)

type StepRouter struct {
	StepController controllers.StepController
}

func NewStepRouter(sc controllers.StepController) StepRouter {
	return StepRouter{
		StepController: sc,
	}
}

func (sr *StepRouter) Register(e *echo.Echo) {
	g := e.Group("/steps")

	g.POST("", sr.StepController.Create)
	g.POST("/:id", sr.StepController.GetById)
}
