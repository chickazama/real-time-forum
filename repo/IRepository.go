package repo

import "matthewhope/real-time-forum/models"

type IRepository interface {
	GetUserByID(id int) (models.User, error)
}
