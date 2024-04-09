package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/jackc/pgx/v5/pgconn"
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
	q := `INSERT INTO ozon.user_profile (id, full_name, email, phone_number, imgsrc) VALUES ($1, $2, $3, $4, $5);`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1=%d, $2=%s, $3=%s, $4=%s, $5=%s", profileRow.ID, profileRow.FullName, profileRow.Email, profileRow.PhoneNumber, profileRow.ImgSrc)))
	start := time.Now()
	_, err := r.storage.Exec(ctx, q, profileRow.ID, profileRow.FullName, profileRow.Email, profileRow.PhoneNumber, profileRow.ImgSrc)
	if err != nil {
		if pgErr := new(pgconn.PgError); errors.As(err, &pgErr) {
			// logs, _ := json.Marshal(pgErr)
			// fmt.Printf("log: \n%s\n", string(logs))
			logger.Error(ctx, fmt.Sprintf("HERE pg err: %v", err))
		} else {
			logger.Error(ctx, fmt.Sprintf("HERE profile already exists: %v", err))
		}
		return 0, models.ErrProfileAlreadyExists
	}
	logger.Info(ctx, fmt.Sprintf("created in %s", time.Since(start)))
	return profileRow.ID, nil
}

func (r *ProfileRepo) GetProfile(ctx context.Context, uID uint) (models.Profile, error) {
	q := `SELECT id, full_name, email, phone_number, imgsrc FROM ozon.user_profile WHERE id = $1;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1=%d", uID)))
	start := time.Now()
	profileRow := &db.ProfileTable{}
	if err := r.storage.Get(ctx, profileRow, q, uID); err != nil {
		if pgErr := new(pgconn.PgError); errors.As(err, &pgErr) {
			// logs, _ := json.Marshal(pgErr)
			// fmt.Printf("log: \n%s\n", string(logs))
			logger.Error(ctx, fmt.Sprintf("HERE pg err: %v", err))
		} else {
			logger.Error(ctx, fmt.Sprintf("HERE profile not found: %v", err))
		}
		return models.Profile{}, models.ErrNoProfile
	}
	logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))

	return db.ConvertTableToProfile(*profileRow), nil
}

func (r *ProfileRepo) UpdateProfile(ctx context.Context, uID uint, profileNew models.Profile) error {
	profileRow := db.ConvertProfileToTable(profileNew)
	q := `UPDATE ozon.user_profile SET full_name=$1, email=$2, phone_number=$3, imgsrc=$4 WHERE id = $5 RETURNING id;`
	start := time.Now()
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1=%s, $2=%s, $3=%s, $4=%s, $5=%d", profileRow.FullName, profileRow.Email, profileRow.PhoneNumber, profileRow.ImgSrc, uID)))
	reqProfileRow := &db.ProfileTable{}
	err := r.storage.Get(ctx, reqProfileRow, q,
		profileNew.FullName,
		profileNew.Email,
		profileNew.PhoneNumber,
		profileNew.ImgSrc,
		uID)
	logger.Info(ctx, fmt.Sprintf("req id is %d", reqProfileRow.ID))
	if err != nil {
		if pgErr := new(pgconn.PgError); errors.As(err, &pgErr) {
			// logs, _ := json.Marshal(pgErr)
			// fmt.Printf("log: \n%s\n", string(logs))
			logger.Error(ctx, fmt.Sprintf("HERE pg err: %v", err))
		} else if reqProfileRow.ID == 0 {
			logger.Error(ctx, fmt.Sprintf("HERE profile not found: %v", err))
		}
		return models.ErrNoProfile
	}
	logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))

	return nil
}

func (r *ProfileRepo) UpdateAvatarByProfileID(ctx context.Context, uID uint, imgSrc string) error {
	q := `UPDATE ozon.user_profile SET imgsrc=$1 WHERE id = $2 RETURNING id;` // Чтобы get ловил отсутствие пользователя
	start := time.Now()
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1=%s, $2=%d", imgSrc, uID)))
	reqProfileRow := &db.ProfileTable{}
	err := r.storage.Get(ctx, reqProfileRow, q,
		imgSrc,
		uID)
	if err != nil {
		if pgErr := new(pgconn.PgError); errors.As(err, &pgErr) {
			logger.Error(ctx, fmt.Sprintf("HERE pg err: %v", err))
		} else if reqProfileRow.ID == 0 {
			logger.Error(ctx, fmt.Sprintf("HERE profile not found: %v", err))
		}
		return models.ErrNoProfile
	}
	logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	return nil
}

func (r *ProfileRepo) GetAvatarByProfileID(ctx context.Context, uID uint) (string, error) {
	q := `SELECT imgsrc, id FROM ozon.user_profile WHERE id = $1;`
	start := time.Now()
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1=%d", uID)))
	reqProfileRow := &db.ProfileTable{}
	err := r.storage.Get(ctx, reqProfileRow, q, uID)
	if err != nil {
		if pgErr := new(pgconn.PgError); errors.As(err, &pgErr) {
			logger.Error(ctx, fmt.Sprintf("HERE pg err: %v", err))
		}
		return "", models.ErrNoProfile
	}

	if reqProfileRow.ID != uID { // Запрос исполнился, но по факту пользователя нет
		logger.Error(ctx, "profile not found")
		return "", models.ErrNoProfile
	}

	logger.Info(ctx, fmt.Sprintf("geted in %s", time.Since(start)))
	return reqProfileRow.ImgSrc, nil
}

func (r *ProfileRepo) DeleteAvatarByProfileID(ctx context.Context, uID uint) (string, error) {
	imgSrc, err := r.GetAvatarByProfileID(ctx, uID)
	if err != nil {
		return "", err
	}
	if imgSrc == "" {
		return "", models.ErrNoAvatar
	}

	q := `UPDATE user_profile SET imgsrc = '' WHERE id = $1;`
	start := time.Now()
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1=%d", uID)))
	_, err = r.storage.Exec(ctx, q, uID)
	if err != nil {
		if pgErr := new(pgconn.PgError); errors.As(err, &pgErr) {
			logger.Error(ctx, fmt.Sprintf("HERE pg err: %v", err))
		}
		return "", models.ErrNoProfile
	}
	logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	return imgSrc, nil
}
