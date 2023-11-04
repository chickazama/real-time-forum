package repo

import (
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/models"
)

type SQLiteRepository struct {
}

func NewSQLiteRepository() *SQLiteRepository {
	return &SQLiteRepository{}
}

func (r *SQLiteRepository) GetUserByID(id int) (models.User, error) {
	return dal.GetUserByID(id)
}

func (r *SQLiteRepository) GetUsers() ([]models.User, error) {
	return dal.GetAllUsers()
}

func (r *SQLiteRepository) GetMessagesBySenderAndTargetIDs(senderID, targetID int) ([]models.Message, error) {
	return dal.GetMessagesBySenderAndTargetIDs(senderID, targetID)
}

func (r *SQLiteRepository) GetPosts() ([]models.Post, error) {
	return dal.GetAllPosts()
}

func (r *SQLiteRepository) GetCommentsByPostID(postID int) ([]models.Comment, error) {
	return dal.GetCommentsByPostID(postID)
}
