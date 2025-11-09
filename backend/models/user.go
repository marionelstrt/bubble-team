package models

import (
	"fmt"

	"math/rand/v2"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int64 `json:"-" db:"id"`

	Boba     string `db:"boba" json:"boba" binding:"required"`
	Name     string `db:"name" json:"name" binding:"required"`
	LastName string `db:"last_name" json:"lastName" binding:"required"`
	Email    string `db:"email" json:"email" binding:"required"`

	EncryptedPassword string `json:"-" db:"encrypted_password"`
	Password          string `json:"password" binding:"required"`

	Verified         bool `json:"-" db:"verified"`
	VerificationCode int  `json:"-" db:"verification_code"`
}

const schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		boba TEXT,
		name TEXT,
		last_name TEXT,
		encrypted_password TEXT,
		email TEXT,
		verified BOOLEAN,
		verification_code INTEGER
	)
`

func CreateUserTable(db *sqlx.DB) {
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

// saves a user in the database from a bound User
func (u *User) Create(db *sqlx.DB) error {
	s, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}
	u.EncryptedPassword = string(s)
	u.VerificationCode = rand.IntN(900000) + 100000
	u.Verified = false

	query := `
		INSERT INTO USERS (
			boba, name, last_name,
			encrypted_password, email,
			verified, verification_code
		) 
		VALUES (
			:boba, :name, :last_name,
			:encrypted_password, :email,
			:verified, :verification_code
		)
		RETURNING id
	`
	var id int64
	rows, err := db.NamedQuery(query, &u)
	if err != nil {
		return fmt.Errorf("failed to insert into users: %w", err)
	}
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to insert into users: %w", err)
	}

	u.ID = id
	return err
}

func (u *User) VerifyCode(db *sqlx.DB, code int) bool {
	if code == u.VerificationCode {
		u.Verified = true
		query := `
			UPDATE users
			SET verified = 1
			WHERE id = :id
		`
		_, err := db.NamedExec(query, &u)
		return err == nil
	}
	return false
}

func (u *User) String() string {
	return fmt.Sprintf("User(Boba: %s, Name: %s, LastName: %s, Password: %s, Email: %s)",
		u.Boba, u.Name, u.LastName, u.Password, u.Email)
}
