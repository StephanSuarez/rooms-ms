package dtos

import "github.com/tesis/internal/rooms/entity"

type RoomResDTO struct {
	ID          string
	Name        string
	Description string
	CreatedBy   string
	NumMaxUsers string
	Status      string
}

func (rrd *RoomResDTO) MapEntityToDto(roomEntity *entity.Room) {
	rrd.ID = roomEntity.ID
	rrd.Name = roomEntity.Name
	rrd.Description = roomEntity.Description
	rrd.CreatedBy = roomEntity.CreatedBy
	rrd.NumMaxUsers = roomEntity.NumMaxUsers
	rrd.Status = roomEntity.Status
}
