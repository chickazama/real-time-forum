package dal

import (
	"matthewhope/real-time-forum/models"
)

func GetAllUsers() ([]models.User, error) {
	var result []models.User
	queryStr := `SELECT * FROM users ORDER BY Nickname ASC;`
	rows, err := identityDb.Query(queryStr)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Nickname, &u.Age, &u.Gender, &u.FirstName, &u.LastName, &u.EmailAddress, &u.EncryptedPassword)
		if err != nil {
			return result, err
		}
		result = append(result, u)
	}
	return result, nil
}
