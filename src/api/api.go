package api

import (
	"github.com/Moji00f/SmsNotification/api/routers"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		routers.NotificationRouter(v1)
	}

	r.Run(":8000")
}
