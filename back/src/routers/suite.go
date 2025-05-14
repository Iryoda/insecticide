package routers

import (
	"github.com/iryoda/insecticide/src/controllers"
	"github.com/labstack/echo/v4"
)

type SuiteRouter struct {
	SuiteController controllers.SuiteController
}

func NewTestRouter(sc controllers.SuiteController) SuiteRouter {
	return SuiteRouter{
		SuiteController: sc,
	}
}

func (sr *SuiteRouter) Register(e *echo.Echo) {
	g := e.Group("/test")

	g.POST("", sr.SuiteController.Create)
	g.POST("/:id", sr.SuiteController.GetById)
}
