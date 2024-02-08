package http

import (
	"github.com/tesis/internal/rooms/repository"
	"github.com/tesis/internal/rooms/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	dbInstance *mongo.Database
	rr         repository.RoomRepository
	rs         services.RoomService
	rh         RoomHandler
}

func (ap *Dependencies) NewAppDependencies(dbInstance *mongo.Database) {
	ap.dbInstance = dbInstance
	ap.rr = repository.NewRoomRepository(dbInstance)
	ap.rs = services.NewRoomService(&ap.rr)
	ap.rh = NewRoomHandler(&ap.rs)
}
