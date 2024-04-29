package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

func (r *AuthRepo) CreateUser(ctx context.Context, user models.User) (uint, error) {
	userRow := dao.ConvertUserToTable(user)
	q := `INSERT INTO default_user (user_login, password_hash) VALUES ($1, $2) returning id;`

	resRow := dao.UserTable{}
	err := r.storage.Get(ctx, &resRow, q, userRow.Login, userRow.PasswordHash)
	if err != nil {
		logger.Info(ctx, "user already exists")
		return 0, fmt.Errorf("user already exists")
	}
	return resRow.ID, nil
}

func (r *AuthRepo) GetUser(ctx context.Context, username string) (models.User, error) {
	q := `SELECT id, password_hash FROM default_user WHERE user_login = $1;`

	userRow := dao.UserTable{}
	if err := r.storage.Get(ctx, &userRow, q, username); err != nil {
		logger.Error(ctx, "user not found")
		return models.User{}, fmt.Errorf("no user found")
	}
	return dao.ConvertTableToUser(userRow), nil
}
