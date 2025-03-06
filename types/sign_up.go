package types

type SignUpDTO struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}
