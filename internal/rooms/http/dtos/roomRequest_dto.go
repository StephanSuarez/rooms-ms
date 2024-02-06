package dtos

import "github.com/tesis/internal/rooms/entity"

type RoomReqDTO struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	CreatedBy   string `json:"userId"`
	NumMaxUsers string `json:"maxUsers"`
	Status      string `json:"status"`
}

func (rrd *RoomReqDTO) MapEntityToDto(roomEntity *entity.Room) {
	rrd.Name = roomEntity.Name
	rrd.Description = roomEntity.Description
	rrd.CreatedBy = roomEntity.CreatedBy
	rrd.NumMaxUsers = roomEntity.NumMaxUsers
	rrd.Status = roomEntity.Status
}

func MapEntityFromDto(RoomReqDTO *RoomReqDTO) *entity.Room {
	return &entity.Room{
		Name:        RoomReqDTO.Name,
		Description: RoomReqDTO.Description,
		CreatedBy:   RoomReqDTO.CreatedBy,
		NumMaxUsers: RoomReqDTO.NumMaxUsers,
		Status:      RoomReqDTO.Status,
	}
}
