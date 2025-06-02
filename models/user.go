package models

import (
	"errors"
	"example.com/RestAPI/db"
	"example.com/RestAPI/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPswd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPswd)
	if err != nil {
		return err
	}
	userid, err := result.LastInsertId()
	u.ID = userid
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHere email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckHashPassword(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
