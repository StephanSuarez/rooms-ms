package services

import (
	"fmt"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository"
)

type roomService struct {
	rr repository.RoomRepository
}

type RoomService interface {
	CreateRoom(roomEntity *entity.Room) error
	GetRooms() ([]entity.Room, error)
	GetRoomByID(id string) (*entity.Room, error)
	UpdateRoom(id string, roomEntity *entity.Room) (*entity.Room, error)
	DeleteRoom(id string) error
}

func NewRoomService(roomRepository *repository.RoomRepository) RoomService {
	return &roomService{
		rr: *roomRepository,
	}
}

func (rs *roomService) CreateRoom(roomEntity *entity.Room) error {

	if !utils.CheckRoomStatus(roomEntity.Status) {
		return fmt.Errorf("room rtatus is not valid")
	}

	if err := rs.rr.InsertOne(roomEntity); err != nil {
		return err
	}

	return nil
}

func (rs *roomService) GetRooms() ([]entity.Room, error) {
	rooms, err := rs.rr.FindAll()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (rs *roomService) GetRoomByID(id string) (*entity.Room, error) {
	room, err := rs.rr.FindOne(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (rs *roomService) UpdateRoom(id string, roomEntity *entity.Room) (*entity.Room, error) {
	roomEntity, err := rs.rr.UpdateOne(id, roomEntity)
	if err != nil {
		return nil, err
	}
	return roomEntity, nil
}

func (rs *roomService) DeleteRoom(id string) error {
	roomDeleted, err := rs.rr.DeleteOne(id)
	if err != nil {
		return err
	}
	if !roomDeleted {
		return fmt.Errorf("room ID was not found")
	}

	return nil
}
