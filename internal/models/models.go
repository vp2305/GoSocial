package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Version   int       `json:"version"`
	Tags      []string  `json:"tags"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentCount int `json:"comments_count"`
}

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

// Not returning the password in the json responses
type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	IsActive  bool     `json:"is_active"`
	CreatedAt string   `json:"created_at"`
	RoleID    int64    `json:"role_id"`
	Role      Role     `json:"role"`
}

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}

type password struct {
	Text *string
	Hash []byte
}

// Helper function
func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	p.Text = &text
	p.Hash = hash

	return nil
}

func (p *password) Compare(payload string) error {
	return bcrypt.CompareHashAndPassword(p.Hash, []byte(payload))
}
