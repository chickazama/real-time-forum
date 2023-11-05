package dal

import "matthewhope/real-time-forum/models"

func GetLimitedMessagesBySenderAndTargetIDs(senderID, targetID, limit, offset int) ([]models.Message, error) {
	var result []models.Message
	queryStr := `SELECT *
	FROM "messages"
	WHERE (SenderID = (?) AND TargetID = (?))
	OR (SenderID = (?) AND TargetID = (?))
	ORDER BY timestamp DESC
	LIMIT (?)
	OFFSET (?);`
	rows, err := forumDb.Query(queryStr, senderID, targetID, targetID, senderID, limit, offset)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.SenderID, &msg.TargetID, &msg.Author, &msg.Content, &msg.Timestamp)
		if err != nil {
			return result, err
		}
		result = append(result, msg)
	}
	return result, nil
}
