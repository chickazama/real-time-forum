package repo

import (
	"matthewhope/real-time-forum/models"
)

type GetUsersRepository interface {
	GetUsers() ([]models.User, error)
}
