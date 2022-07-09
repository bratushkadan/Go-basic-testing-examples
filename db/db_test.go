package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type DBMock struct {
	emails map[string]bool
}

func (db *DBMock) UserExists(email string) bool {
	return db.emails[email]
}

func (db *DBMock) AddUser(email string) {
	db.emails[email] = true
}

func TestNewUser(t *testing.T) {
	errPattern := `user with '%s' email already exists`

	table := []struct {
		name    string
		email   string
		present bool
		wanterr bool
	}{
		{`success`, `gregorysmith@myexampledomain.com`, false, false},
		{`error`, `johndoe@myexampledomain.com`, true, true},
	}

	for _, item := range table {
		t.Run(item.name, func(t *testing.T) {
			// создаём объек-заглушку
			dbMock := &DBMock{emails: make(map[string]bool)}
			if item.present {
				dbMock.AddUser(item.email)
			}
			// выполняем наш код, передавая объект-заглушку
			err := NewUser(dbMock, item.email)
			if !item.wanterr {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, fmt.Sprintf(errPattern, item.email))
			}
		})
	}
}
