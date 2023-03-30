package response

// Captcha 验证码
type Captcha struct {
	B64s          string `json:"picPath"`
	CaptchaId     string `json:"captchaId"`
	OpenCaptcha   bool   `json:"openCaptcha"`
	CaptchaLength int    `json:"captchaLength"`
}
