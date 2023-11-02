package models

type Post struct {
	ID         int    `json:"id"`
	AuthorID   int    `json:"authorID"`
	Author     string `json:"author"`
	Content    string `json:"content"`
	Categories string `json:"categories"`
	Timestamp  int64  `json:"timestamp"`
}
