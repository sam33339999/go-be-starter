package domains

type AuthService interface {
	// middleware check token valid
	Authorize(token string) (bool, error)
	// for mock login
	Attempt(id uint) (string, error)
	// create jwt token
	CreateToken() (string, string)
	// refresh jwt token
	RefreshToken(refreshToken string) (string, string, error)
}
