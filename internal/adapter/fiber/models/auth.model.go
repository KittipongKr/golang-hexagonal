package fiber_models

type AuthReqb struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
