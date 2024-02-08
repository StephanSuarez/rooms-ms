package repository

import (
	"fmt"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type roomRepository struct {
	dbInstance *mongo.Database
}

type RoomRepository interface {
	FindAll()
	InsertOne(roomEntity *entity.Room) error
}

func NewRoomRepository(dbMongoInstance *mongo.Database) RoomRepository {
	return &roomRepository{
		dbInstance: dbMongoInstance,
	}
}

// func CreateRoom(dbInstance *mongo.Database, room *models.Room) error {
// 	collection := dbInstance.Collection("rooms")
// 	ctx := context.TODO()

// 	_, err := collection.InsertOne(ctx, room)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (rr *roomRepository) FindAll() {}

func (rr *roomRepository) InsertOne(roomEntity *entity.Room) error {
	roomModel := models.Room{}
	roomModel.MapEntityToModel(roomEntity)
	if !utils.CheckRoomStatus(roomModel.Status) {
		return fmt.Errorf("room status is no valid")
	}
	return nil
}
