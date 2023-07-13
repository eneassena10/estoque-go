package entities

type User struct {
	ID       int
	Name     string
	Nickname string
	Password string
	Logado   int
}

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) WithID(id int) *User {
	u.ID = id
	return u
}

func (u *User) WithName(name string) *User {
	u.Name = name
	return u
}

func (u *User) WithNickname(nickname string) *User {
	u.Nickname = nickname
	return u
}

func (u *User) WithPassword(password string) *User {
	u.Password = password
	return u
}

func (u *User) WithLogado(logado int) *User {
	u.Logado = logado
	return u
}
