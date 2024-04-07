package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
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
	profileRow := db.ConvertProfileToTable(profile)
	q := `INSERT INTO user_profile (id, 
                          full_name, 
                          email, 
                          phone_number, 
                          imgsrc
                          ) VALUES ($1, $2, $3, $4, $5);`
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("created in %s", time.Since(start)))
	}()
	_, err := r.storage.Exec(ctx, q, profileRow)
	if err != nil {
		logger.Error(ctx, "profile already exists")
		return 0, models.ErrUserAlreadyExists
	}
	return profileRow.ID, nil
}

func (r *ProfileRepo) GetProfile(ctx context.Context, uID uint) (models.Profile, error) {
	q := `SELECT * FROM default_user WHERE id = $1;`
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	}()
	profileRow := &db.ProfileTable{}
	if err := r.storage.Get(ctx, profileRow, q, uID); err != nil {
		logger.Error(ctx, "profile not found")
		return models.Profile{}, models.ErrNoUser
	}
	return db.ConvertTableToProfile(*profileRow), nil
}

func (r *ProfileRepo) UpdateProfile(ctx context.Context, uID uint, profileNew models.Profile) error {
	q := `UPDATE user_profile SET email=$1, full_name=$2, phone_number=$3, imgscr=$4 WHERE id = $6`
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	}()
	_, err := r.storage.Exec(ctx, q,
		profileNew.Email,
		profileNew.FullName,
		profileNew.PhoneNumber,
		profileNew.ImgSrc,
		uID)
	if err != nil {
		logger.Error(ctx, "profile not found")
		return models.ErrNoProfile
	}

	return nil
}
