package auth

type UseCase interface {
	SignIn(credentials Credentials) error
	SignUp(credentials Credentials) (string, error)
}
