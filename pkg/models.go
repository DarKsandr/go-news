package pkg

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	About     string    `json:"about"`
	AvatarUrl string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type News struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ImageUrl   string    `json:"image_url"`
	ReadTime   int       `json:"read_time"`
	Views      int       `json:"views"`
	ShareCount int       `json:"share_count"`
	UserID     int       `json:"user_id"`
	User       User      `json:"user"`
	CategoryID int       `json:"category_id"`
	Category   Category  `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type NewsApi struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	ImageUrl   string      `json:"image_url"`
	ReadTime   int         `json:"read_time"`
	Views      int         `json:"views"`
	ShareCount int         `json:"share_count"`
	UserID     int         `json:"user_id"`
	User       User        `json:"user"`
	CategoryID int         `json:"category_id"`
	Category   CategoryApi `json:"category"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	News []News `json:"news"`
}

type CategoryApi struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CategoryApi) TableName() string {
	return "categories"
}
