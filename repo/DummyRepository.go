package repo

import (
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/models"
)

type DummyRepository struct {
}

func NewDummyRepository() *DummyRepository {
	return &DummyRepository{}
}

func (r *DummyRepository) GetUserByID(id int) (models.User, error) {
	return dal.GetUserByID(id)
}

func (r *DummyRepository) GetUsers() ([]models.User, error) {
	return dal.GetAllUsers()
}

func (r *DummyRepository) GetMessagesBySenderAndTargetIDs(senderID, targetID int) ([]models.Message, error) {
	return dal.GetMessagesBySenderAndTargetIDs(senderID, targetID)
}

func (r *DummyRepository) GetPosts() ([]models.Post, error) {
	return dal.GetAllPosts()
}

func (r *DummyRepository) GetCommentsByPostID(postID int) ([]models.Comment, error) {
	return dal.GetCommentsByPostID(postID)
}
