package service

type JwtService interface {
	Encode(username string) string
	Verify()
}
