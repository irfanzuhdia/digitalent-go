package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/orders", createOrder)
	router.GET("/orders", getOrders)
	router.PATCH("/orders/:id", updateOrder)
	router.DELETE("/orders/:id", deleteOrder)
}
