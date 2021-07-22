package postgres

import (
	"database/sql"

	"github.com/benbjohnson/myapp"
	_ "github.com/lig/pq"
)

type UserService struct {
	DB                 *sql.DB
	TransactionService myapp.TransactionService
}

func (s *UserService) User(id int) (*myapp.User, error) {
	var u myapp.User
	row := s.DB.QueryRow(`select id, name from users where id = $1`, id)
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) CreateUser(u *myapp.User) error {
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	return nil
}

func (s *UserService) Users() ([]*myapp.User, error) {
	return nil, nil
}

func Open(addr string) (*sql.DB, error) {
	return nil, nil
}
