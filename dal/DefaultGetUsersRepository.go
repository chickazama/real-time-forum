package dal

import (
	"database/sql"
	"matthewhope/real-time-forum/models"
)

type DefaultGetUsersRepository struct {
	db *sql.DB
}

func NewDefaultGetUsersRepository() *DefaultGetUsersRepository {
	return &DefaultGetUsersRepository{db: identityDb}
}

func (r *DefaultGetUsersRepository) GetUsers() ([]models.User, error) {
	return GetAllUsers(r.db)
}
