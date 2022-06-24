package request

type Login struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}

type GetName struct {
	Username string `json:"username"`
	Forced   bool   `json:"forced"`
}

type NewPassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
