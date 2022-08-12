package dto

// Login credential
type LoginCredentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type ProfileRequestType struct {
	ID        int64  `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Token     string `json:"token"`
	ImageUrl  string `json:"url"`
}

type ScreechRequestType struct {
	ID            int64  `json:"id"`
	Content       string `json:"content"`
	CreatorID     int64  `json:"user_id"`
	OrderByAscend bool   `json:"order_by_ascend"`
}
