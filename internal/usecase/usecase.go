package usecase

type Auth interface {
	Login(login string, password string) (string, error)
	Signup(login string, password string) (string, string, error)
	Logout(login string) error
}
