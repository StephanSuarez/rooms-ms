package repository

import (
	"context"
	"fmt"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type roomRepository struct {
	dbInstance *mongo.Database
}

var collection *mongo.Collection

type RoomRepository interface {
	InsertOne(roomEntity *entity.Room) error
	FindAll() ([]entity.Room, error)
	FindOne(id string) (*entity.Room, error)
	UpdateOne(id string, roomEntity *entity.Room) (*entity.Room, error)
	DeleteOne(id string) (bool, error)
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

func (rr *roomRepository) FindOne(id string) (*entity.Room, error) {
	var room models.Room

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&room)
	if err != nil {
		return nil, err
	}

	return room.MapEntityFromModel(), nil
}

func (rr *roomRepository) UpdateOne(id string, roomEntity *entity.Room) (*entity.Room, error) {
	room := models.Room{}
	room.MapEntityToModel(roomEntity)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	update := bson.M{
		"$set": room,
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("the document to update was not found")
	}

	roomEntity.ID = id

	return roomEntity, nil
}

func (rr *roomRepository) DeleteOne(id string) (bool, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		fmt.Print(err)
		return false, nil
	}

	fmt.Println(result)

	return false, nil
}
