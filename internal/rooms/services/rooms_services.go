package services

import (
	"fmt"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository"
)

type roomService struct {
	rr *repository.RoomRepository
}

type RoomService interface {
	FindAll()
	CreateRoom(roomEntity *entity.Room) error
}

func NewRoomService(roomRepository *repository.RoomRepository) RoomService {
	return &roomService{
		rr: roomRepository,
	}
}

func (rs *roomService) FindAll() {
	fmt.Println("hello")
}

func (rs *roomService) CreateRoom(roomEntity *entity.Room) error {
	fmt.Println("service")
	fmt.Println(roomEntity)
	if !utils.CheckRoomStatus(roomEntity.Status) {
		return fmt.Errorf("room rtatus is not valid")
	}

	return nil
}
