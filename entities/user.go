package entities

import "context"

type User struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	RealName   string `json:"realName"`
	Address    string `json:"address"`
	LoginToken string `json:"loginToken"`
}

type UserUsecase interface {
	Register(ctx context.Context, username string, password string) (token string, err error)
	Login(ctx context.Context, username string, password string) (token string, err error)
	Logout(ctx context.Context, token string) (err error)
	Authorize(ctx context.Context, token string) (user User, err error)
}

type UserRepository interface {
	Create(ctx context.Context, username string, password string, loginToken string) (err error)
	Get(ctx context.Context, username string, password string) (user User, err error)
	GetByToken(ctx context.Context, token string) (user User, err error)
	UpdateToken(ctx context.Context, username string, loginToken string) (err error)
	IsUsernameExist(ctx context.Context, username string) bool
}
