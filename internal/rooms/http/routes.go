package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	routesRooms := r.Group("/v1/rooms/")
	routesRooms.GET("/", )
	routesRooms.POST("/", )

	r.Run()
}
