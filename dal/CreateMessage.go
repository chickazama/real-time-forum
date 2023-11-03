package dal

func CreateMessage(senderID, targetID int, author, content string, timestamp int64) (int, error) {
	var result int
	queryStr := `INSERT INTO "messages" (SenderID, TargetID, Author, Content, Timestamp) VALUES (?, ?, ?, ?, ?);`
	stmt, err := forumDb.Prepare(queryStr)
	if err != nil {
		return result, err
	}
	res, err := stmt.Exec(senderID, targetID, author, content, timestamp)
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
