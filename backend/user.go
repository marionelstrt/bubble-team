package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Boba              string `db:"boba" json:"boba" binding:"required"`
	Name              string `db:"name" json:"name" binding:"required"`
	LastName          string `db:"last_name" json:"lastName" binding:"required" `
	Password          string `json:"password" binding:"required"`
	Email             string `db:"email" json:"email" binding:"required"`
	EncryptedPassword string `json:"-" db:"encrypted_password"`
}

const schema = `
	CREATE TABLE IF NOT EXISTS users (
		id integer PRIMARY KEY AUTOINCREMENT,
		boba text,
		name text,
		last_name text,
		encrypted_password text,
		email text
	)
`

// creates the users table inside the database
func CreateTable(db *sqlx.DB) {
	db.MustExec(schema)
}

func FindUserByID(db *sqlx.DB, id int) *User {
	var user User
	row := db.QueryRow("SELECT * FROM users where id=?", id)
	err := row.Scan(&user)
	if err != nil {
		return nil
	}
	return &user
}

func Encrypt(password []byte) string {
	s, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	s1 := string(s)
	return s1
}

// saves a user in the database from a bound User
func (u *User) Save(db *sqlx.DB) (*User, error) {
	s, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt password: %w", err)
	}
	u.EncryptedPassword = string(s)
	query := `
		INSERT INTO USERS (boba, name, last_name, encrypted_password, email) 
		VALUES (:boba, :name, :last_name, :encrypted_password, :email)
	`
	_, err = db.NamedExec(query, &u)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into users: %w", err)
	}
	return u, err
}

func (u *User) PrettyPrint() string {
	return fmt.Sprintf("User(Boba: %s, Name: %s, LastName: %s, Password: %s, Email: %s)",
		u.Boba, u.Name, u.LastName, u.Password, u.Email)
}
