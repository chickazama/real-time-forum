package dal

import (
	"database/sql"
	"matthewhope/real-time-forum/models"
)

type DummyRepository struct {
	db *sql.DB
}

func NewDummyRepository() *DummyRepository {
	return &DummyRepository{db: identityDb}
}

func (r *DummyRepository) GetUserByID(id int) (models.User, error) {
	return GetUserByID(id)
}
