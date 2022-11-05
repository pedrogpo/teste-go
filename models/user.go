package models

type User struct {
	nome  string
	idade int
}

func NewUser() *User {
	user := User{
		nome:  "Rafael",
		idade: 22,
	}

	return &user
}

func (u *User) SetNome(nome string) {
	u.nome = nome
}

func (u User) GetNome() string {
	return u.nome
}
