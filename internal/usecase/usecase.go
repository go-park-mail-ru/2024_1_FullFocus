package usecase

type Auth interface {
	Login(login string, password string) (string, error)
	Signup(login string, password string) (string, string, error)
	GetUserIDBySessionID(sID string) (uint, error)
	Logout(sID string) error
	IsLoggedIn(isID string) bool
}
