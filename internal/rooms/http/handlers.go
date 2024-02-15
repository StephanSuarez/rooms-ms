package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/http/dtos"
	"github.com/tesis/internal/rooms/services"
)

type roomHandler struct {
	rs services.RoomService
}

type RoomHandler interface {
	CreateRoom(ctx *gin.Context)
	GetRooms(ctx *gin.Context)
}

func NewRoomHandler(roomService *services.RoomService) RoomHandler {
	return &roomHandler{
		rs: *roomService,
	}
}

func (rh *roomHandler) CreateRoom(ctx *gin.Context) {
	roomDto := dtos.RoomReqDTO{}

	if err := ctx.ShouldBind(&roomDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Body",
		})
		return
	}

	if !utils.CheckRoomStatus(roomDto.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Room Status is not valid",
		})
		return
	}

	if err := rh.rs.CreateRoom(roomDto.MapEntityFromDto()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": roomDto,
	})
}

func (rh *roomHandler) GetRooms(ctx *gin.Context) {
	rooms, err := rh.rs.GetRooms()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var roomsResDto []dtos.RoomResDTO
	for i := 0; i < len(rooms); i++ {
		roomsResDto = append(roomsResDto, dtos.RoomResDTO(rooms[i]))
	}

	ctx.JSON(http.StatusOK, roomsResDto)
}
