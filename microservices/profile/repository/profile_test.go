package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mockdb "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/repository/dao"
)

func TestNewRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mockdb.NewMockDatabase(ctrl)
	defer ctrl.Finish()
	t.Run("Check Repo creation", func(t *testing.T) {
		pr := repository.NewRepo(db)
		require.NotEmpty(t, pr, "profile repo not created")
	})
}

func TestCreateProfile(t *testing.T) {
	testCases := []struct {
		name          string
		profile       models.Profile
		mockBehavior  func(*mockdb.MockDatabase, string, uint)
		expectedError error
	}{
		{
			name: "Test successful profile creation",
			profile: models.Profile{
				ID:         1,
				FullName:   "test",
				AvatarName: "aaa",
			},
			mockBehavior: func(d *mockdb.MockDatabase, q string, id uint) {
				d.EXPECT().Exec(context.Background(), q, id).Return(mockdb.MockSQLResult{}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Test duplicate profile creation",
			profile: models.Profile{
				ID:         1,
				FullName:   "test",
				AvatarName: "aaa",
			},
			mockBehavior: func(d *mockdb.MockDatabase, q string, id uint) {
				d.EXPECT().Exec(context.Background(), q, id).Return(mockdb.MockSQLResult{}, sql.ErrNoRows)
			},
			expectedError: models.ErrProfileAlreadyExists,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mockdb.NewMockDatabase(ctrl)
			defer ctrl.Finish()
			testCase.mockBehavior(db, "INSERT INTO user_profile (id) VALUES (?);", testCase.profile.ID)
			pr := repository.NewRepo(db)

			err := pr.CreateProfile(context.Background(), testCase.profile.ID)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestGetProfile(t *testing.T) {
	testCases := []struct {
		name          string
		id            uint
		mockBehavior  func(*mockdb.MockDatabase, *dao.ProfileTable, string, uint)
		expectedError error
	}{
		{
			name: "Test successful get",
			id:   1,
			mockBehavior: func(d *mockdb.MockDatabase, u *dao.ProfileTable, q string, id uint) {
				d.EXPECT().Get(context.Background(), u, q, id).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test not existing get",
			id:   1,
			mockBehavior: func(d *mockdb.MockDatabase, u *dao.ProfileTable, q string, id uint) {
				d.EXPECT().Get(context.Background(), u, q, id).Return(sql.ErrNoRows)
			},
			expectedError: models.ErrNoProfile,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := mockdb.NewMockDatabase(ctrl)
			q := `SELECT id, full_name, address, phone_number, gender, imgsrc
	FROM user_profile
	WHERE id = ?;`

			testCase.mockBehavior(db, &dao.ProfileTable{}, q, testCase.id)
			pr := repository.NewRepo(db)
			_, err := pr.GetProfile(context.Background(), testCase.id)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

// TODO: fix
func TestUpdateProfile(t *testing.T) {
	testCases := []struct {
		name          string
		profile       models.Profile
		mockBehavior  func(d *mockdb.MockDatabase, q string, name, address string, gender uint, id uint)
		expectedError error
	}{
		{
			name: "Test successful get",
			profile: models.Profile{
				ID:         1,
				FullName:   "test",
				AvatarName: "aaa",
			},
			mockBehavior: func(d *mockdb.MockDatabase, q string, name, address string, gender uint, id uint) {
				d.EXPECT().Exec(context.Background(), q, name, address, gender, id).Return(mockdb.MockSQLResult{}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Test fail get",
			profile: models.Profile{
				ID:         1,
				FullName:   "test",
				AvatarName: "aaa",
			},
			mockBehavior: func(d *mockdb.MockDatabase, q string, name, address string, gender uint, id uint) {
				d.EXPECT().Exec(context.Background(), q, name, address, gender, id).Return(mockdb.MockSQLResult{}, sql.ErrNoRows)
			},
			expectedError: models.ErrNoProfile,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := mockdb.NewMockDatabase(ctrl)

			testCase.mockBehavior(db,
				`UPDATE user_profile
		SET full_name = ?, address = ?, gender = ?
		WHERE id = ?
		RETURNING id;`, testCase.profile.FullName, testCase.profile.FullName, 0, testCase.profile.ID)
			pr := repository.NewRepo(db)
			err := pr.UpdateProfile(context.Background(), testCase.profile.ID, models.ProfileUpdateInput{
				FullName: testCase.profile.FullName,
			})
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestGetProfileNamesByIDs(t *testing.T) {
	testCases := []struct {
		name           string
		ids            []uint
		mockBehavior   func(d *mockdb.MockDatabase, names *[]string, q string, pIDs []uint, pIDsOrder []uint)
		expectedResult []string
		expectedError  error
	}{
		{
			name: "Test 0 id passed",
			ids:  []uint{},
			mockBehavior: func(d *mockdb.MockDatabase, names *[]string, q string, pIDs []uint, pIDsOrder []uint) {
				d.EXPECT().Select(context.Background(), names, q, pIDs, pIDsOrder).Return(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  models.ErrNoProfile,
		},
		{
			name: "Test not all profiles found",
			ids:  []uint{23, 7, 35, 4},
			mockBehavior: func(d *mockdb.MockDatabase, names *[]string, q string, pIDs []uint, pIDsOrder []uint) {
				d.EXPECT().Select(context.Background(), names, q, pIDs, pIDsOrder).
					SetArg(1, []string{"i", "love", "mail.ru"}).Return(nil)
			},
			expectedResult: nil,
			expectedError:  models.ErrNoProfile,
		},
		{
			name: "Test no profiles found",
			ids:  []uint{1, 2, 3},
			mockBehavior: func(d *mockdb.MockDatabase, names *[]string, q string, pIDs []uint, pIDsOrder []uint) {
				d.EXPECT().Select(context.Background(), names, q, pIDs, pIDsOrder).Return(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  models.ErrNoProfile,
		},
		{
			name: "Test all profiles found",
			ids:  []uint{1, 2, 3},
			mockBehavior: func(d *mockdb.MockDatabase, names *[]string, q string, pIDs []uint, pIDsOrder []uint) {
				d.EXPECT().Select(context.Background(), names, q, pIDs, pIDsOrder).
					SetArg(1, []string{"i", "love", "mail.ru"}).Return(nil)
			},
			expectedResult: []string{"i", "love", "mail.ru"},
			expectedError:  nil,
		},
		{
			name: "Test more than needed profiles found",
			ids:  []uint{1, 2},
			mockBehavior: func(d *mockdb.MockDatabase, names *[]string, q string, pIDs []uint, pIDsOrder []uint) {
				d.EXPECT().Select(context.Background(), names, q, pIDs, pIDsOrder).
					SetArg(1, []string{"i", "love", "mail.ru"}).Return(nil)
			},
			expectedResult: nil,
			expectedError:  models.ErrNoProfile,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := mockdb.NewMockDatabase(ctrl)
			q := `SELECT full_name
	FROM user_profile
	WHERE id = ANY (?)
	ORDER BY array_position(?, id);`

			var names []string
			testCase.mockBehavior(db, &names, q, testCase.ids, testCase.ids)
			pr := repository.NewRepo(db)
			result, err := pr.GetProfileNamesByIDs(context.Background(), testCase.ids)
			require.Equal(t, testCase.expectedResult, result, "Wrong result")
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestGetProfileMetaInfo(t *testing.T) {
	testCases := []struct {
		name           string
		pID            uint
		mockBehavior   func(d *mockdb.MockDatabase, info *dao.ProfileMetaInfo, q string, pID uint)
		expectedResult models.ProfileMetaInfo
		expectedError  error
	}{
		{
			name: "Test no profile found",
			pID:  0,
			mockBehavior: func(d *mockdb.MockDatabase, info *dao.ProfileMetaInfo, q string, pID uint) {
				d.EXPECT().Get(context.Background(), info, q, pID).Return(sql.ErrNoRows)
			},
			expectedResult: models.ProfileMetaInfo{},
			expectedError:  models.ErrNoProfile,
		},
		{
			name: "Test found",
			pID:  1,
			mockBehavior: func(d *mockdb.MockDatabase, info *dao.ProfileMetaInfo, q string, pID uint) {
				d.EXPECT().Get(context.Background(), info, q, pID).
					SetArg(1, dao.ProfileMetaInfo{
						FullName:   "tester",
						AvatarName: "avatar",
					}).Return(nil)
			},
			expectedResult: models.ProfileMetaInfo{
				FullName:   "tester",
				AvatarName: "avatar",
			},
			expectedError: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := mockdb.NewMockDatabase(ctrl)
			var info dao.ProfileMetaInfo
			testCase.mockBehavior(db, &info, "SELECT full_name, imgsrc FROM user_profile WHERE id = ?;", testCase.pID)
			pr := repository.NewRepo(db)
			result, err := pr.GetProfileMetaInfo(context.Background(), testCase.pID)
			require.Equal(t, testCase.expectedResult, result, "Wrong info")
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestGetAvatarByProfileID(t *testing.T) {
	testCases := []struct {
		name           string
		pID            uint
		mockBehavior   func(d *mockdb.MockDatabase, avatar *string, q string, pID uint)
		expectedResult string
		expectedError  error
	}{
		{
			name: "Test no profile found",
			pID:  0,
			mockBehavior: func(d *mockdb.MockDatabase, avatar *string, q string, pID uint) {
				d.EXPECT().Get(context.Background(), avatar, q, pID).Return(sql.ErrNoRows)
			},
			expectedResult: "",
			expectedError:  models.ErrNoProfile,
		},
		{
			name: "Test found",
			pID:  1,
			mockBehavior: func(d *mockdb.MockDatabase, avatar *string, q string, pID uint) {
				d.EXPECT().Get(context.Background(), avatar, q, pID).
					SetArg(1, "avatar").Return(nil)
			},
			expectedResult: "avatar",
			expectedError:  nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := mockdb.NewMockDatabase(ctrl)
			q := `SELECT imgsrc
	FROM user_profile
	WHERE id = ?;`

			var avatar string
			testCase.mockBehavior(db, &avatar, q, testCase.pID)
			pr := repository.NewRepo(db)
			result, err := pr.GetAvatarByProfileID(context.Background(), testCase.pID)
			require.Equal(t, testCase.expectedResult, result, "Wrong info")
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestUpdateAvatarByProfileID(t *testing.T) {
	testCases := []struct {
		name           string
		pID            uint
		avatarName     string
		mockBehavior   func(d *mockdb.MockDatabase, prevAvatar *string, q string, avatarName string, pID uint)
		expectedResult string
		expectedError  error
	}{
		{
			name:       "Test no profile found",
			pID:        0,
			avatarName: "new",
			mockBehavior: func(d *mockdb.MockDatabase, prevAvatar *string, q string, avatarName string, pID uint) {
				d.EXPECT().Get(context.Background(), prevAvatar, q, pID, avatarName, pID).Return(sql.ErrNoRows)
			},
			expectedResult: "",
			expectedError:  models.ErrNoProfile,
		},
		{
			name:       "Test found",
			pID:        1,
			avatarName: "new",
			mockBehavior: func(d *mockdb.MockDatabase, prevAvatar *string, q string, avatarName string, pID uint) {
				d.EXPECT().Get(context.Background(), prevAvatar, q, pID, avatarName, pID).
					SetArg(1, "someName").Return(nil)
			},
			expectedResult: "someName",
			expectedError:  nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := mockdb.NewMockDatabase(ctrl)
			var avatar string
			testCase.mockBehavior(db, &avatar, `WITH prev_imgsrc AS (
    		  SELECT imgsrc
    		  FROM user_profile
    		  WHERE id = ?
		  )
		  UPDATE user_profile
		  SET imgsrc = ?
		  WHERE id = ?
		  RETURNING (SELECT imgsrc FROM prev_imgsrc);`, testCase.avatarName, testCase.pID)
			pr := repository.NewRepo(db)
			result, err := pr.UpdateAvatarByProfileID(context.Background(), testCase.pID, testCase.avatarName)
			require.Equal(t, testCase.expectedResult, result, "Wrong info")
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestDeleteAvatarByProfileID(t *testing.T) {
	testCases := []struct {
		name           string
		pID            uint
		mockBehavior   func(d *mockdb.MockDatabase, avatar *string, q string, pID uint)
		expectedResult string
		expectedError  error
	}{
		{
			name: "Test no profile found",
			pID:  0,
			mockBehavior: func(d *mockdb.MockDatabase, avatar *string, q string, pID uint) {
				d.EXPECT().Get(context.Background(), avatar, q, pID, pID).Return(sql.ErrNoRows)
			},
			expectedResult: "",
			expectedError:  models.ErrNoProfile,
		},
		{
			name: "Test found",
			pID:  1,
			mockBehavior: func(d *mockdb.MockDatabase, avatar *string, q string, pID uint) {
				d.EXPECT().Get(context.Background(), avatar, q, pID, pID).
					SetArg(1, "avatar").Return(nil)
			},
			expectedResult: "avatar",
			expectedError:  nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := mockdb.NewMockDatabase(ctrl)
			var avatar string
			testCase.mockBehavior(db, &avatar, `WITH prev_imgsrc AS (
    	  	  SELECT imgsrc
    	  	  FROM user_profile
    	  	  WHERE id = ?
		  )
		  UPDATE user_profile
	  	  SET imgsrc = ''
		  WHERE id = ?
		  RETURNING (SELECT imgsrc FROM prev_imgsrc);`, testCase.pID)
			pr := repository.NewRepo(db)
			result, err := pr.DeleteAvatarByProfileID(context.Background(), testCase.pID)
			require.Equal(t, testCase.expectedResult, result, "Wrong info")
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}
