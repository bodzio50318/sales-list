package api

type loginBody struct {
	UserName string
	Password string
}

type jwtRespone struct {
	UserName string
	JwtToken string
}
