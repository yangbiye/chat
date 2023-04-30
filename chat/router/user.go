package router

import (
	"github.com/chat/api"
	"github.com/chat/service"
	"github.com/gin-gonic/gin"
)

func RouterSetup() {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.POST("user/register", api.UserRegister)
		v1.GET("ws", service.WsHandler)
	}
	r.Run(":8080")
}
