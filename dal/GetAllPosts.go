package dal

import "matthewhope/real-time-forum/models"

func GetAllPosts() ([]models.Post, error) {
	var result []models.Post
	queryStr := `SELECT * FROM posts ORDER BY Timestamp DESC;`
	rows, err := forumDb.Query(queryStr)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.ID, &p.AuthorID, &p.Author, &p.Content, &p.Categories, &p.Timestamp)
		if err != nil {
			return result, err
		}
		result = append(result, p)
	}
	return result, nil
}
