package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repo struct {
	storage db.Database
}

func NewRepo(dbClient db.Database) *Repo {
	return &Repo{
		storage: dbClient,
	}
}

func (r *Repo) CreateProfile(ctx context.Context, pID uint) error {
	q := `INSERT INTO user_profile (id) VALUES (?);`

	_, err := r.storage.Exec(ctx, q, pID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.ErrProfileAlreadyExists
	}
	return nil
}

func (r *Repo) GetProfile(ctx context.Context, uID uint) (models.Profile, error) {
	q := `SELECT id, full_name, address, phone_number, gender, imgsrc
	FROM user_profile
	WHERE id = ?;`

	var profileRow dao.ProfileTable
	if err := r.storage.Get(ctx, &profileRow, q, uID); err != nil {
		logger.Error(ctx, err.Error())
		return models.Profile{}, models.ErrNoProfile
	}
	return dao.ConvertTableToProfile(profileRow), nil
}

func (r *Repo) GetProfileNamesByIDs(ctx context.Context, pIDs []uint) ([]string, error) {
	q := `SELECT full_name
	FROM user_profile
	WHERE id = ANY (?)
	ORDER BY array_position(?, id);`

	var names []string
	if err := r.storage.Select(ctx, &names, q, pIDs, pIDs); err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrNoProfile
	}
	if len(pIDs) != len(names) {
		return nil, models.ErrNoProfile
	}
	return names, nil
}

func (r *Repo) GetProfileMetaInfo(ctx context.Context, pID uint) (models.ProfileMetaInfo, error) {
	q := `SELECT full_name, imgsrc
	FROM user_profile
	WHERE id = ?;`

	var info dao.ProfileMetaInfo
	if err := r.storage.Get(ctx, &info, q, pID); err != nil {
		logger.Error(ctx, err.Error())
		return models.ProfileMetaInfo{}, models.ErrNoProfile
	}
	return dao.ConvertProfileMetaInfo(info), nil
}

func (r *Repo) GetProfileNamesAvatarsByIDs(ctx context.Context, pIDs []uint) ([]models.ProfileNameAvatar, error) {
	q := `SELECT full_name, imgsrc
	FROM user_profile
	WHERE id = ANY(?)
	ORDER BY array_position(?, id);`

	profileData := make([]dao.ProfileNameAvatar, 0)
	if err := r.storage.Select(ctx, &profileData, q, pIDs, pIDs); err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrInternal
	}
	if len(pIDs) != len(profileData) {
		return nil, models.ErrNoProfile
	}
	return dao.ConvertProfileNamesAvatarsToModels(profileData), nil
}

func (r *Repo) UpdateProfile(ctx context.Context, uID uint, profileNew models.ProfileUpdateInput) error {
	q := `UPDATE user_profile
	SET full_name = ?, address = ?, gender = ? %s
	WHERE id = ?
	RETURNING id;`

	var err error
	if profileNew.PhoneNumber != "" {
		q = fmt.Sprintf(q, "phone_num = ?")
		_, err = r.storage.Exec(ctx, q, profileNew.FullName, profileNew.Address, profileNew.Gender, profileNew.PhoneNumber, uID)
	} else {
		q := fmt.Sprintf(q, "")
		_, err = r.storage.Exec(ctx, q, profileNew.FullName, profileNew.Address, profileNew.Gender, uID)
	}
	if err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		var sqlErr *pgconn.PgError
		if errors.As(err, &sqlErr) {
			if strings.Contains(sqlErr.Message, "unique") {
				return models.ErrPhoneAlreadyExists
			}
			return models.ErrNoProfile
		}
	}
	return nil
}

func (r *Repo) UpdateAvatarByProfileID(ctx context.Context, uID uint, imgSrc string) (string, error) {
	q := `WITH prev_imgsrc AS (
    		  SELECT imgsrc
    		  FROM user_profile
    		  WHERE id = ?
		  )
		  UPDATE user_profile
		  SET imgsrc = ?
		  WHERE id = ?
		  RETURNING (SELECT imgsrc FROM prev_imgsrc);`

	var prevImgSrc string
	if err := r.storage.Get(ctx, &prevImgSrc, q, uID, imgSrc, uID); err != nil {
		logger.Error(ctx, err.Error())
		return "", models.ErrNoProfile
	}
	return prevImgSrc, nil
}

func (r *Repo) GetAvatarByProfileID(ctx context.Context, uID uint) (string, error) {
	q := `SELECT imgsrc
	FROM user_profile
	WHERE id = ?;`

	var imgSrc string
	err := r.storage.Get(ctx, &imgSrc, q, uID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return "", models.ErrNoProfile
	}
	return imgSrc, nil
}

func (r *Repo) DeleteAvatarByProfileID(ctx context.Context, uID uint) (string, error) {
	q := `WITH prev_imgsrc AS (
    	  	  SELECT imgsrc
    	  	  FROM user_profile
    	  	  WHERE id = ?
		  )
		  UPDATE user_profile
	  	  SET imgsrc = ''
		  WHERE id = ?
		  RETURNING (SELECT imgsrc FROM prev_imgsrc);`

	var prevImgSrc string
	if err := r.storage.Get(ctx, &prevImgSrc, q, uID, uID); err != nil {
		logger.Error(ctx, err.Error())
		return "", models.ErrNoProfile
	}
	return prevImgSrc, nil
}
