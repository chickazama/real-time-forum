package repo

import "matthewhope/real-time-forum/models"

type IRepository interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
}
