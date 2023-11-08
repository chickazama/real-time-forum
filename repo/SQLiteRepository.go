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

func (r *SQLiteRepository) CreateComment(postID, authorID int, author, content string, timestamp int64) (int, error) {
	return dal.CreateComment(postID, authorID, author, content, timestamp)
}
func (r *SQLiteRepository) CreateUser(nickname, age, gender, firstName, lastName, emailAddress, password string) (int, error) {
	return dal.CreateUser(nickname, age, gender, firstName, lastName, emailAddress, password)
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

func (r *SQLiteRepository) GetLimitedMessagesBySenderAndTargetIDs(senderID, targetID, limit, offset int) ([]models.Message, error) {
	return dal.GetLimitedMessagesBySenderAndTargetIDs(senderID, targetID, limit, offset)
}

func (r *SQLiteRepository) GetPosts() ([]models.Post, error) {
	return dal.GetAllPosts()
}

func (r *SQLiteRepository) GetCommentsByPostID(postID int) ([]models.Comment, error) {
	return dal.GetCommentsByPostID(postID)
}
