//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/wanglixianyii/go-middleware/go-es/common"
	"github.com/wanglixianyii/go-middleware/go-es/config"
	"github.com/wanglixianyii/go-middleware/go-es/dao"
	"github.com/wanglixianyii/go-middleware/go-es/handler"
	"github.com/wanglixianyii/go-middleware/go-es/service"
)

func InitializeHandler(conf *config.ServerConfig) *handler.UserHandler {
	wire.Build(
		common.NewEsClient,
		common.NewRouterClient,

		dao.NewUserES,
		service.NewUserService,

		handler.NewUserHandler,
	)
	return &handler.UserHandler{}
}
