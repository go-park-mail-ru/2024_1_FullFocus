package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

func (r *AuthRepo) CreateUser(ctx context.Context, user models.User) (uint, error) {
	userRow := dao.ConvertUserToTable(user)
	q := `INSERT INTO default_user (email, password_hash) VALUES ($1, $2) returning id;`

	resRow := dao.UserTable{}
	err := r.storage.Get(ctx, &resRow, q, userRow.Email, userRow.PasswordHash)
	if err != nil {
		logger.Error(ctx, err.Error())
		return 0, models.ErrUserAlreadyExists
	}
	return resRow.ID, nil
}

func (r *AuthRepo) GetUser(ctx context.Context, email string) (models.User, error) {
	q := `SELECT id, password_hash FROM default_user WHERE email = $1;`

	userRow := dao.UserTable{}
	if err := r.storage.Get(ctx, &userRow, q, email); err != nil {
		logger.Error(ctx, err.Error())
		return models.User{}, models.ErrUserNotFound
	}
	return dao.ConvertTableToUser(userRow), nil
}

func (r *AuthRepo) GetUserEmailByUserID(ctx context.Context, uID uint) (string, error) {
	q := `SELECT email FROM default_user WHERE id = $1;`

	var email string
	if err := r.storage.Get(ctx, &email, q, uID); err != nil {
		logger.Error(ctx, err.Error())
		return "", models.ErrUserNotFound
	}
	return email, nil
}

func (r *AuthRepo) GetUserPassword(ctx context.Context, userID uint) (string, error) {
	q := `SELECT password_hash FROM default_user WHERE id = ?;`

	var password string
	err := r.storage.Get(ctx, &password, q, userID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return "", models.ErrUserNotFound
	}
	return password, nil
}

func (r *AuthRepo) UpdatePassword(ctx context.Context, userID uint, password string) error {
	q := `UPDATE default_user SET password_hash = ? WHERE id = ?;`

	_, err := r.storage.Exec(ctx, q, password, userID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.ErrUserNotFound
	}
	return nil
}
