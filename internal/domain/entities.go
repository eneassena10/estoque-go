package domain

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type IUser interface {
	Logar(u User) error
	Logout(u User) error
	Create(u User) (User, error)
}
