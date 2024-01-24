package repositories

type roomStatus string

const (
	active   roomStatus = "Active"
	inactive roomStatus = "Inactive"
	deleted  roomStatus = "Deleted"
)

type Room struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"desc"`
	CreatedBy   string     `json:"userId"`
	NumMaxUsers string     `json:"maxUsers"`
	Status      roomStatus `json:"status"` // active, desactive, deleted
}

func FindAll() *[]Room {
	rooms := []Room{
		{
			ID:          "123",
			Name:        "roomTesis",
			Description: "desc1",
			CreatedBy:   "stephanId",
			NumMaxUsers: "0",
			Status:      active,
		},
		{
			ID:          "456",
			Name:        "roomTesis002",
			Description: "desc2",
			CreatedBy:   "stephanId",
			NumMaxUsers: "5",
			Status:      inactive,
		},
	}

	return &rooms
}

func stringPtr(s string) *string {
	return &s
}
