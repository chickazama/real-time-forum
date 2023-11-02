package dal

import "golang.org/x/crypto/bcrypt"

func CreateUser(nickname, age, gender, firstName, lastName, emailAddress, password string) error {
	queryStr := `INSERT INTO "users" (Nickname, Age, Gender, FirstName, LastName, EmailAddress, EncryptedPassword) VALUES(?, ?, ?, ?, ?, ?, ?);`
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	stmt, err := identityDb.Prepare(queryStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(nickname, age, gender, firstName, lastName, emailAddress, string(encryptedPassword))
	if err != nil {
		return err
	}
	return nil
}
