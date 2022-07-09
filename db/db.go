package db

import "fmt"

type DB struct {
	// some fields intrinsic to a database
}

type DBStorage interface {
	UserExists(email string) bool
}

func (db *DB) UserExists(email string) bool {

	// let's assume there's some complex network logic
	// ...

	return true
}

func NewUser(db DBStorage, email string) error {
	if db.UserExists(email) {
		return fmt.Errorf(`user with '%s' email already exists`, email)
	}
	// добавляем запись
	return nil
}
