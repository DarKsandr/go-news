package pkg

type NewsRequest struct {
	Title      string `json:"title" validate:"required,alpha,min=3"`
	Content    string `json:"content" validate:"required,alpha,min=3"`
	ImageUrl   string `json:"image_url" validate:"required,uri"`
	ReadTime   int    `json:"read_time" validate:"required,number"`
	Views      int    `json:"views" validate:"required,number"`
	ShareCount int    `json:"share_count" validate:"required,number"`
	UserID     int    `json:"user_id" validate:"required,number"`
	CategoryID int    `json:"category_id" validate:"required,number"`
}
