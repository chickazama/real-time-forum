package dal

import "matthewhope/real-time-forum/models"

func GetUserByNickname(nickname string) (models.User, error) {
	var result models.User
	queryStr := `SELECT * FROM "users" WHERE Nickname = (?);`
	row := identityDb.QueryRow(queryStr, nickname)
	err := row.Scan(&result.ID, &result.Nickname, &result.Age, &result.Gender, &result.FirstName, &result.LastName, &result.EmailAddress, &result.EncryptedPassword)
	if err != nil {
		return result, err
	}
	return result, nil
}
