package dal

import "golang.org/x/crypto/bcrypt"

func CreateUser(nickname, age, gender, firstName, lastName, emailAddress, password string) (int, error) {
	var result int
	queryStr := `INSERT INTO "users" (Nickname, Age, Gender, FirstName, LastName, EmailAddress, EncryptedPassword) VALUES(?, ?, ?, ?, ?, ?, ?);`
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return result, err
	}
	stmt, err := identityDb.Prepare(queryStr)
	if err != nil {
		return result, err
	}
	res, err := stmt.Exec(nickname, age, gender, firstName, lastName, emailAddress, string(encryptedPassword))
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
