package models

import "github.com/tesis/internal/rooms/entity"

type RoomRes struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"desc"`
	CreatedBy   string `bson:"userId"`
	NumMaxUsers string `bson:"maxUsers"`
	Status      string `bson:"status"` // active, desactive, deleted
}

func (model *RoomRes) MapEntityToModel(roomEntity *entity.Room) {
	model.Name = roomEntity.Name
	model.Description = roomEntity.Description
	model.CreatedBy = roomEntity.CreatedBy
	model.NumMaxUsers = roomEntity.NumMaxUsers
	model.Status = roomEntity.Status
}

func (model *RoomRes) MapEntityFromModel() *entity.Room {
	return &entity.Room{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedBy:   model.CreatedBy,
		NumMaxUsers: model.NumMaxUsers,
		Status:      string(model.Status),
	}
}
