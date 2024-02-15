package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, appRoomDependencies *RoomDependencies) {
	routesRooms := r.Group("/v1/rooms/")
	routesRooms.GET("/", appRoomDependencies.rh.GetRooms)
	routesRooms.POST("/", appRoomDependencies.rh.CreateRoom)

	r.Run()
}
