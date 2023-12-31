package dal

import (
	"matthewhope/real-time-forum/models"
)

func GetUserByID(id int) (models.User, error) {
	var result models.User
	queryStr := `SELECT * FROM "users" WHERE ID = (?);`
	row := identityDb.QueryRow(queryStr, id)
	err := row.Scan(&result.ID, &result.Nickname, &result.Age, &result.Gender, &result.FirstName, &result.LastName, &result.EmailAddress, &result.EncryptedPassword)
	if err != nil {
		return result, err
	}
	return result, nil
}
