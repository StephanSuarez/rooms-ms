package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/v1/rooms/find-all", GetRooms)

	r.Run()
}
