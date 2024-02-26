package usecase

type AuthUsecase struct {
	repo AuthRepo
}

type AuthRepo interface {
	Login(login string, password string) (string, error)
	Signup(login string, password string) (string, uint, error)
	IsLoggedIn(login string) (bool, error)
}

func (uc *AuthUsecase) Login(login string, password string) (string, error) {
	return uc.repo.Login(login, password)
}

func (uc *AuthUsecase) Signup(login string, password string) (string, uint, error) {
	return uc.repo.Signup(login, password)
}

func (uc *AuthUsecase) IsLoggedIn(login string) (bool, error) {
	return uc.repo.IsLoggedIn(login)
}
