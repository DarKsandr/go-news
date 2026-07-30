package pkg

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	About     string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type News struct {
	ID         int
	Title      string
	Content    string
	ImageUrl   string
	ReadTime   int
	Views      int
	ShareCount int
	UserID     int
	User       User
	CategoryID int
	Category   Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Category struct {
	ID   int
	Name string
	News []News
}
