package server

import (
	"golang_tugas_3/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	orders *controllers.OrderControllers
}

func NewRouter(orders *controllers.OrderControllers) *Router {
	return &Router{orders: orders}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	router.POST("/orders", r.orders.POST_Orders)
	router.GET("/orders", r.orders.GET_Orders)
	router.PUT("/orders/:order_id", r.orders.PUT_Orders)
	router.DELETE("/orders/:order_id", r.orders.DELETE_Orders)

	router.Run(port)
}
