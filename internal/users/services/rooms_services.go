package services

import "github.com/tesis/internal/users/repositories"

func GetRooms() *[]repositories.Room {
	// business logic
	return repositories.FindAll()
}
