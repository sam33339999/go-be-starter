package domains

type AuthService interface {
	Authorize(token string) (bool, error)
	CreateToken() string
}
