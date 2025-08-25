package entity

type Users struct {
	ID       int    `json:"id_user"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
