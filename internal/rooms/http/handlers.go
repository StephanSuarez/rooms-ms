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
	GetRoomByID(ctx *gin.Context)
	UpdateRoom(ctx *gin.Context)
	DeleteRoom(ctx *gin.Context)
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

func (rh *roomHandler) GetRoomByID(ctx *gin.Context) {
	id := ctx.Param("id")
	room, err := rh.rs.GetRoomByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	roomResDto := dtos.RoomResDTO{}
	roomResDto.MapEntityToDto(room)

	ctx.JSON(http.StatusOK, roomResDto)
}

func (rh *roomHandler) UpdateRoom(ctx *gin.Context) {
	id := ctx.Param("id")

	roomDto := dtos.RoomReqDTO{}
	if err := ctx.ShouldBind(&roomDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if !utils.CheckRoomStatus(roomDto.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Room Status is not valid",
		})
		return
	}

	roomEntity, err := rh.rs.UpdateRoom(id, roomDto.MapEntityFromDto())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roomEntity)
}

func (rh *roomHandler) DeleteRoom(ctx *gin.Context) {

}
