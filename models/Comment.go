package models

type Comment struct {
	ID         int    `json:"id"`
	PostID     int    `json:"postID"`
	AuthorID   int    `json:"authorID"`
	Author     string `json:"author"`
	Content    string `json:"content"`
	Categories string `json:"categories"`
	Timestamp  int64  `json:"timestamp"`
}
