package models

import "github.com/tesis/internal/rooms/entity"

type roomStatus string

const (
	active   roomStatus = "Active"
	inactive roomStatus = "Inactive"
	deleted  roomStatus = "Deleted"
)

type Room struct {
	ID          string     `bson:"id"`
	Name        string     `bson:"name"`
	Description string     `bson:"desc"`
	CreatedBy   string     `bson:"userId"`
	NumMaxUsers string     `bson:"maxUsers"`
	Status      roomStatus `bson:"status"` // active, desactive, deleted
}

func (model *Room) MapEntityToModel(roomEntity *entity.Room) {
	model.Name = roomEntity.Name
	model.Description = roomEntity.Description
	model.CreatedBy = roomEntity.CreatedBy
	model.NumMaxUsers = roomEntity.NumMaxUsers
	model.Status = roomStatus(roomEntity.Status)
}

func MapEntityFromModel() {}
