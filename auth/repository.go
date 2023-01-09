package auth

type Repository interface {
	Get(username string) (string, error)
	Insert(username, password string) (string, error)
}
