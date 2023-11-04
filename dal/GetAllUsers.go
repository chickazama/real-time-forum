package dal

import (
	"database/sql"
	"matthewhope/real-time-forum/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	var result []models.User
	queryStr := `SELECT * FROM users ORDER BY Nickname ASC;`
	rows, err := db.Query(queryStr)
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
