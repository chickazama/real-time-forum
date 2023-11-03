package dal

import "matthewhope/real-time-forum/models"

func GetCommentsByPostID(postID int) ([]models.Comment, error) {
	var result []models.Comment
	queryStr := `SELECT * FROM "comments" WHERE PostID = (?);`
	rows, err := forumDb.Query(queryStr, postID)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var c models.Comment
		err = rows.Scan(&c.ID, &c.PostID, &c.AuthorID, &c.Author, &c.Content, &c.Timestamp)
		if err != nil {
			return result, err
		}
		result = append(result, c)
	}
	return result, nil
}
