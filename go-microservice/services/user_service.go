package services

import (
	"sync"

	"go-microservice/models"
)

type UserService struct {
	mu     sync.Mutex
	users  map[int]models.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (s *UserService) Create(user models.User) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user
	return user
}

func (s *UserService) GetAll() []models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := []models.User{}
	for _, u := range s.users {
		result = append(result, u)
	}
	return result
}

func (s *UserService) GetByID(id int) (models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[id]
	return u, ok
}

func (s *UserService) Update(id int, user models.User) (models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return models.User{}, false
	}
	user.ID = id
	s.users[id] = user
	return user, true
}

func (s *UserService) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return false
	}
	delete(s.users, id)
	return true
}
