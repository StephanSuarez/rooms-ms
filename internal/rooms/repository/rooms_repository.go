package repository

import (
	// "context"

	"fmt"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository/models"
	// "go.mongodb.org/mongo-driver/mongo"
)

// func (rm *Room) MapToRoom(roomDto http.RoomDTO) *Room {
// 	rm.ID = roomDto.ID
// 	rm.Name = roomDto.Name
// 	rm.Description = roomDto.Description
// 	rm.CreatedBy = roomDto.CreatedBy
// 	rm.NumMaxUsers = roomDto.NumMaxUsers
// 	rm.Status = rm.Status
// 	return rm
// }

// type roomRepository struct {
// 	dbInstance *mongo.Database
// }

// type RoomRepository interface {
// 	CreateRoom(room *Room) error
// }

// func NewRoomRepository(dbMongoInstance *mongo.Database) RoomRepository {
// 	return &roomRepository{
// 		dbInstance: dbMongoInstance,
// 	}
// }

// func CreateRoom(dbInstance *mongo.Database, room *models.Room) error {
// 	collection := dbInstance.Collection("rooms")
// 	ctx := context.TODO()

// 	_, err := collection.InsertOne(ctx, room)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func FindAll() {}

func InsertOne(roomEntity *entity.Room) error {
	roomModel := models.Room{}
	roomModel.MapEntityToModel(roomEntity)
	if !utils.CheckRoomStatus(roomModel.Status) {
		return fmt.Errorf("room status is no valid")
	}
	return nil
}
