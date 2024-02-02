package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/rooms/services"
)

func GetRooms(ctx *gin.Context) {
	services.FindAll()

	ctx.JSON(http.StatusOK, gin.H{
		"rooms": "rooms",
	})
}
