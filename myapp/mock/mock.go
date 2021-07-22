package mock

import "github.com/benbjohnson/myapp"

type UserService struct {
	UserFn      func(id int) (*myapp.User, error)
	UserInvoked bool

	UsersFn      func() ([]*myapp.User, error)
	UsersInvoked bool
}

func (s *UserService) User(id int) (*myapp.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
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
