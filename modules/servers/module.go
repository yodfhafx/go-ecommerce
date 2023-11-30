package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yodfhafx/go-ecommerce/modules/middlewares/middlewaresHandlers"
	"github.com/yodfhafx/go-ecommerce/modules/middlewares/middlewaresRepositories"
	"github.com/yodfhafx/go-ecommerce/modules/middlewares/middlewaresUsecases"
	"github.com/yodfhafx/go-ecommerce/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecases(repository)
	return middlewaresHandlers.MiddlewaresHandlers(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}
