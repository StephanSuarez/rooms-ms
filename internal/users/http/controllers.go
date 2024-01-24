package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/users/services"
)

func GetRooms(ctx *gin.Context) {
	rooms := services.GetRooms()
	if len(*rooms) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "there are not rooms",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"rooms": *rooms,
	})
}
