package user

import (
	"context"
	"errors"
	"regexp"
	"sync"
)

var (
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidID    = errors.New("invalid id")
	ErrNotFound     = errors.New("not found")
	emailRE         = regexp.MustCompile(`^\S+@\S+\.\S+$`)
)

type User struct {
	Name  string
	Email string
	ID    string
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrInvalidName
	}
	if !emailRE.MatchString(u.Email) {
		return ErrInvalidEmail
	}
	if u.ID == "" {
		return ErrInvalidID
	}
	return nil
}

// UserManager manages users
// Contains a map of users, a mutex, and a context

type UserManager struct {
	ctx   context.Context
	users map[string]User // userID -> User
	mutex sync.RWMutex    // Protects users map
}

// NewUserManager creates a new UserManager
func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[string]User),
	}
}

// NewUserManagerWithContext creates a new UserManager with context
func NewUserManagerWithContext(ctx context.Context) *UserManager {
	return &UserManager{
		ctx:   ctx,
		users: make(map[string]User),
	}
}

// AddUser adds a user
func (m *UserManager) AddUser(u User) error {
	if m.ctx != nil && m.ctx.Err() != nil {
		return m.ctx.Err()
	}
	if err := u.Validate(); err != nil {
		return err
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.users[u.ID] = u
	return nil
}

// RemoveUser removes a user
func (m *UserManager) RemoveUser(id string) error {
	if m.ctx != nil && m.ctx.Err() != nil {
		return m.ctx.Err()
	}

	m.mutex.Lock()
	m.mutex.Unlock()

	delete(m.users, id)
	return nil
}

// GetUser retrieves a user by id
func (m *UserManager) GetUser(id string) (User, error) {
	if m.ctx != nil && m.ctx.Err() != nil {
		return User{}, m.ctx.Err()
	}

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	u, ok := m.users[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}
