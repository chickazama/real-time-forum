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
	return GetUserByID(r.db, id)
}

func (r *DummyRepository) GetUsers() ([]models.User, error) {
	return GetAllUsers(r.db)
}

func (r *DummyRepository) GetMessagesBySenderAndTargetIDs(senderID, targetID int) ([]models.Message, error) {
	return GetMessagesBySenderAndTargetIDs(senderID, targetID)
}

func (r *DummyRepository) GetPosts() ([]models.Post, error) {
	return GetAllPosts()
}
