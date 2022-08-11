package mysql

import (
	"context"

	"gorm.io/gorm"

	"car-record/entities"
)

func (user) TableName() string {
	return "user"
}

type user struct {
	ID         uint   `gorm:"primaryKey,column:id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	RealName   string `gorm:"column:real_name"`
	Address    string `gorm:"column:address"`
	LoginToken string `gorm:"column:login_token"`
	// CreateTime time.Time `gorm:"autoCreateTime,column:create_time"`
}

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) entities.UserRepository {
	return &mysqlUserRepository{Conn}
}

func (r *mysqlUserRepository) Create(ctx context.Context, username string, password string, loginToken string) (err error) {
	newUser := user{Username: username, Password: password, LoginToken: loginToken}
	err = r.Conn.WithContext(ctx).Create(&newUser).Error

	return
}

func (r *mysqlUserRepository) Get(ctx context.Context, username string, password string) (entities.User, error) {
	queryUser := user{Username: username, Password: password}
	result := r.Conn.WithContext(ctx).Model(&user{}).Where(&queryUser).First(&queryUser)

	return entities.User{
		ID:         queryUser.ID,
		Username:   queryUser.Username,
		RealName:   queryUser.RealName,
		Address:    queryUser.Address,
		LoginToken: queryUser.LoginToken,
	}, result.Error
}

func (r *mysqlUserRepository) GetByToken(ctx context.Context, token string) (entities.User, error) {
	queryUser := user{LoginToken: token}
	result := r.Conn.WithContext(ctx).Model(&user{}).Where(&queryUser).First(&queryUser)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return entities.User{
		ID:         queryUser.ID,
		Username:   queryUser.Username,
		RealName:   queryUser.RealName,
		Address:    queryUser.Address,
		LoginToken: queryUser.LoginToken,
	}, nil
}

func (r *mysqlUserRepository) UpdateToken(ctx context.Context, username string, loginToken string) (err error) {
	err = r.Conn.WithContext(ctx).Model(&user{}).Where(&user{Username: username}).Update("login_token", loginToken).Error

	return
}

func (r *mysqlUserRepository) IsUsernameExist(ctx context.Context, username string) bool {
	var count int64
	r.Conn.WithContext(ctx).Model(&user{}).Where(&user{Username: username}).Count(&count)

	return count > 0
}
