//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/wanglixianyii/go-middleware/go-wire/dao"
	"github.com/wanglixianyii/go-middleware/go-wire/handler"
	"github.com/wanglixianyii/go-middleware/go-wire/router"
	"github.com/wanglixianyii/go-middleware/go-wire/service"
)

func InitializeHandler() *handler.UserHandler {
	wire.Build(
		router.NewRouterClient,
		dao.NewUserModel,
		service.NewUserService,
		handler.NewUserHandler,
	)
	return &handler.UserHandler{}
}
