package models

import (
	"errors"

	"chaitanyaallu.dev/event-management/db"
	"chaitanyaallu.dev/event-management/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) CreateUser() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmnt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func ValidateCredentials(email, password string) error {
	user, err := GetUserByEmail(email)
	if err != nil {
		return err
	}
	isPasswordCorrect := utils.ComparePasswords(user.Password, password)
	if !isPasswordCorrect {
		return errors.New("invalid credentials")
	}
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	query := "SELECT id, email, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)
	user := User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id int64) (*User, error) {
	query := "SELECT id, email, password FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	user := User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
