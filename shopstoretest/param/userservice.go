package param

type UserRegisterRequest struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Credit      int    `json:"credit"`
}

type UserRegisterResponse struct {
	UserInfo UserInfo `json:"user_info"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	UserInfo UserInfo `json:"user"`
	Tokens   Tokens   `json:"tokens"`
}

type UserProfileRequest struct {
	ID uint `json:"id"`
}

type UserProfileResponse struct {
	UserInfo UserInfo `json:"user_info"`
}
