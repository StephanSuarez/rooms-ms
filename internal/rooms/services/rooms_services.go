package services

import (
	"fmt"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository"
)

func FindAll() {
	fmt.Println("hello")
}

func CreateRoom(roomEntity *entity.Room) error {
	fmt.Println("service")
	fmt.Println(roomEntity)
	if !utils.CheckRoomStatus(roomEntity.Status) {
		return fmt.Errorf("room rtatus is not valid")
	}
	repository.InsertOne(roomEntity)
	return nil
}

