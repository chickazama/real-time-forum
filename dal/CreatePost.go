package dal

func CreatePost(authorID int, author, content, categories string, timestamp int64) (int, error) {
	var result int
	queryStr := `INSERT INTO "posts" (AuthorID, Author, Content, Categories, Timestamp) VALUES(?, ?, ?, ?, ?);`
	stmt, err := forumDb.Prepare(queryStr)
	if err != nil {
		return result, err
	}
	res, err := stmt.Exec(authorID, author, content, categories, timestamp)
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
