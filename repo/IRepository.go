package repo

import "matthewhope/real-time-forum/models"

type IRepository interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	GetMessagesBySenderAndTargetIDs(userID, targetID int) ([]models.Message, error)
	GetLimitedMessagesBySenderAndTargetIDs(senderID, targetID, limit, offset int) ([]models.Message, error)
	GetPosts() ([]models.Post, error)
	GetCommentsByPostID(postID int) ([]models.Comment, error)
}
