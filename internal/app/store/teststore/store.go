package teststore

import (
	"github.com/DmitryZzz/http-rest-api/internal/app/store"
	"github.com/DmitryZzz/http-rest-api/internal/app/model"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository	
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
