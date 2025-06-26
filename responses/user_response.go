package responses

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
