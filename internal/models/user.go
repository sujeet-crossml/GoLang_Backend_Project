package models

import "github.com/sujeet-crossml/GoLang_Backend_Project/internal/config"

// defining User model
type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"` // never return password in json
}

func CreateUser(user *User) error {
	query := "INSERT INTO users (name, email, hash_password) VALUES (?, ?, ?)"
	res, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return  err
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)
	return nil
}

func GetUserByEmail(email string) (*User, error){
	u := &User{}
	query := "SELECT id, name, email, hash_password FROM users WHERE email = ?"
	err := config.DB.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByID(id int) (*User, error) {
	u := &User{}
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := config.DB.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}