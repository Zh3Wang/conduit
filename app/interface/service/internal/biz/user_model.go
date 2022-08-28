package biz

type UpdateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
