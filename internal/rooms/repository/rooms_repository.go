package repository

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

func FindAll() {}
