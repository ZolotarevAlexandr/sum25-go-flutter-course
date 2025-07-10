package userdomain

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidName     = errors.New("invalid name")
)

var (
	emailRE     = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	upperCaseRE = regexp.MustCompile(`[A-Z]`)
	lowerCaseRE = regexp.MustCompile(`[a-z]`)
	numberRE    = regexp.MustCompile(`[0-9]`)
)

// User represents a user entity in the domain
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"-"` // Never serialize password
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser creates a new user with validation
// Requirements:
// - Email must be valid format
// - Name must be 2-51 characters
// - Password must be at least 8 characters
// - CreatedAt and UpdatedAt should be set to current time
func NewUser(email, name, password string) (*User, error) {
	user := &User{
		ID:        0,
		Email:     email,
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	// Check email, name, and password validity
	if err := ValidateEmail(u.Email); err != nil {
		return err
	}
	if err := ValidateName(u.Name); err != nil {
		return err
	}
	if err := ValidatePassword(u.Password); err != nil {
		return err
	}
	return nil
}

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) error {
	// Use regex pattern to validate email format
	// Email should not be empty and should match standard email pattern
	if !emailRE.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

// ValidateName checks if name is valid
func ValidateName(name string) error {
	// Name should be 2-50 characters, trimmed of whitespace
	// Should not be empty after trimming
	name = strings.TrimSpace(name)
	if len(name) < 2 || len(name) > 50 {
		return ErrInvalidName
	}
	return nil
}

// ValidatePassword checks if password meets security requirements
func ValidatePassword(password string) error {
	// Password should be at least 8 characters
	// Should contain at least one uppercase, lowercase, and number
	if !upperCaseRE.MatchString(password) || !lowerCaseRE.MatchString(password) || !numberRE.MatchString(password) {
		return ErrInvalidPassword
	}
	if len(password) < 8 {
		return ErrInvalidPassword
	}
	return nil
}

// UpdateName updates the user's name with validation
func (u *User) UpdateName(name string) error {
	if err := ValidateName(name); err != nil {
		return err
	}
	u.Name = strings.TrimSpace(name)
	u.UpdatedAt = time.Now()
	return nil
}

// UpdateEmail updates the user's email with validation
func (u *User) UpdateEmail(email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	if err := ValidateEmail(email); err != nil {
		return err
	}
	u.Email = strings.ToLower(strings.TrimSpace(email))
	u.UpdatedAt = time.Now()
	return nil
}
