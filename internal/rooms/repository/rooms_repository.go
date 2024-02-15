package repository

import (
	"context"
	"fmt"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type roomRepository struct {
	dbInstance *mongo.Database
}

var collection *mongo.Collection

type RoomRepository interface {
	InsertOne(roomEntity *entity.Room) error
	FindAll() ([]entity.Room, error)
}

func NewRoomRepository(dbMongoInstance *mongo.Database) RoomRepository {
	collection = dbMongoInstance.Collection("rooms")
	return &roomRepository{
		dbInstance: dbMongoInstance,
	}
}

func (rr *roomRepository) InsertOne(roomEntity *entity.Room) error {
	roomModel := models.Room{}
	roomModel.MapEntityToModel(roomEntity)
	fmt.Print("inserting ")

	if !utils.CheckRoomStatus(roomModel.Status) {
		return fmt.Errorf("room status is no valid")
	}

	ctx := context.TODO()
	_, err := collection.InsertOne(ctx, roomModel)
	fmt.Print("inserting ")
	if err != nil {
		panic(err)
	}
	return nil
}

func (rr *roomRepository) FindAll() ([]entity.Room, error) {
	var rooms []models.RoomRes
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &rooms)
	if err != nil {
		return nil, err
	}

	roomsEntity := []entity.Room{}
	for i := 0; i < len(rooms); i++ {
		roomentity := rooms[i].MapEntityFromModel()
		roomsEntity = append(roomsEntity, *roomentity)
	}

	return roomsEntity, nil
}
