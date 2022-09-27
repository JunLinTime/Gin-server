package request

type Register struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

type Login struct {
	Username  string `json:"username"` // 用户名
	Password  string `json:"password"` // 密码
	Captcha   string `json:"captcha"`  // 验证码
	CaptchaId string `json:"captchaId"`
}
