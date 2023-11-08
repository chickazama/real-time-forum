package repo

import "matthewhope/real-time-forum/models"

type IRepository interface {
	CreateComment(postID, authorID int, author, content string, timestamp int64) (int, error)
	CreateMessage(senderID, targetID int, author, content string, timestamp int64) (int, error)
	CreatePost(authorID int, author, content, categories string, timestamp int64) (int, error)
	CreateUser(nickname, age, gender, firstName, lastName, emailAddress, password string) (int, error)
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	GetMessagesBySenderAndTargetIDs(userID, targetID int) ([]models.Message, error)
	GetLimitedMessagesBySenderAndTargetIDs(senderID, targetID, limit, offset int) ([]models.Message, error)
	GetPosts() ([]models.Post, error)
	GetCommentsByPostID(postID int) ([]models.Comment, error)
}
