package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wanglixianyii/go-middleware/go-wire/service"
	"log"
	"net/http"
)

type UserHandler struct {
	engine      *gin.Engine
	userService *service.UserService
}

func NewUserHandler(engine *gin.Engine, userService *service.UserService) *UserHandler {
	return &UserHandler{
		engine:      engine,
		userService: userService,
	}
}

func (h *UserHandler) GetName() string {
	return h.userService.GetName()
}

func (h *UserHandler) Run() {

	gin.ForceConsoleColor()
	h.engine.Use(gin.Logger())
	h.engine.Use(gin.Recovery())
	h.registerRouter()

	err := h.engine.Run()
	if err != nil {
		log.Fatalln("server start failed")
	}
}

func (h *UserHandler) registerRouter() {
	u := h.engine.Group("api/user")
	{
		u.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "hello"})
		})

	}
}
