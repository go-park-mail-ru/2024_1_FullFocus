package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
)

type ProfileRepo struct {
	storage db.Database
}

func NewProfileRepo(dbClient db.Database) *ProfileRepo {
	return &ProfileRepo{
		storage: dbClient,
	}
}

func (r *ProfileRepo) CreateProfile(ctx context.Context, profile models.Profile) (uint, error) {
	q := `INSERT INTO ozon.user_profile (id, full_name, email, phone_number, imgsrc) VALUES ($1, $2, $3, $4, $5);`
	_, err := r.storage.Exec(ctx, q, profile.ID, profile.FullName, profile.Email, profile.PhoneNumber, profile.ImgSrc)
	if err != nil {
		logger.Error(ctx, "insert error: "+err.Error())
		return 0, models.ErrProfileAlreadyExists
	}
	return profile.ID, nil
}

func (r *ProfileRepo) GetProfile(ctx context.Context, uID uint) (models.Profile, error) {
	q := `SELECT id, full_name, email, phone_number, imgsrc FROM ozon.user_profile WHERE id = $1;`
	var profileRow dao.ProfileTable
	if err := r.storage.Get(ctx, &profileRow, q, uID); err != nil {
		logger.Error(ctx, "select error: "+err.Error())
		return models.Profile{}, models.ErrNoProfile
	}
	return dao.ConvertTableToProfile(profileRow), nil
}

func (r *ProfileRepo) UpdateProfile(ctx context.Context, uID uint, profileNew models.Profile) error {
	q := `UPDATE ozon.user_profile SET full_name=$1, email=$2, phone_number=$3, imgsrc=$4 WHERE id = $5 RETURNING id;`
	_, err := r.storage.Exec(ctx, q,
		profileNew.FullName,
		profileNew.Email,
		profileNew.PhoneNumber,
		profileNew.ImgSrc,
		uID)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("profile update error: "+err.Error()))
		return models.ErrNoProfile
	}
	return nil
}

func (r *ProfileRepo) UpdateAvatarByProfileID(ctx context.Context, uID uint, imgSrc string) error {
	q := `UPDATE ozon.user_profile SET imgsrc=$1 WHERE id = $2;`
	_, err := r.storage.Exec(ctx, q, imgSrc, uID)
	if err != nil {
		logger.Error(ctx, "update error: "+err.Error())
		return models.ErrNoProfile
	}
	return nil
}

func (r *ProfileRepo) GetAvatarByProfileID(ctx context.Context, uID uint) (string, error) {
	q := `SELECT imgsrc FROM ozon.user_profile WHERE id = $1;`
	var imgSrc string
	err := r.storage.Get(ctx, &imgSrc, q, uID)
	if err != nil {
		logger.Error(ctx, "profile select error: "+err.Error())
		return "", models.ErrNoProfile
	}
	return imgSrc, nil
}

func (r *ProfileRepo) DeleteAvatarByProfileID(ctx context.Context, uID uint) error {
	q := `UPDATE user_profile SET imgsrc = '' WHERE id = $1;`
	_, err := r.storage.Exec(ctx, q, uID)
	if err != nil {
		logger.Error(ctx, "update error: "+err.Error())
		return models.ErrNoProfile
	}
	return nil
}
