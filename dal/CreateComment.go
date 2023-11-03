package dal

func CreateComment(postID, authorID int, author, content string, timestamp int64) (int, error) {
	var result int
	queryStr := `INSERT INTO "comments" (PostID, AuthorID, Author, Content, Timestamp) VALUES (?, ?, ?, ?, ?);`
	stmt, err := forumDb.Prepare(queryStr)
	if err != nil {
		return result, err
	}
	res, err := stmt.Exec(postID, authorID, author, content, timestamp)
	if err != nil {
		return result, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return result, err
	}
	result = int(id)
	return result, nil
}
