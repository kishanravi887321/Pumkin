package requests

type UserRegisterRequest struct {

	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserLoginByGoogleRequest struct {
	GoogleToken string `json:"google_token" binding:"required"`
}