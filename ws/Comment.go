package ws

type Comment struct {
	Code int         `json:"code"`
	Data CommentData `json:"data"`
}
type CommentData struct {
	ID        int    `json:"id"`
	PostID    int    `json:"postID"`
	AuthorID  int    `json:"authorID"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}
