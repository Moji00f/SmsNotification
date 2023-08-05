package routers

import (
	"github.com/Moji00f/SmsNotification/api/handlers"

	"github.com/gin-gonic/gin"
)

func NotificationRouter(r *gin.RouterGroup) {
	h := handlers.NewNotificationHandler()
	r.POST("/sendalert", h.Alert)
}
