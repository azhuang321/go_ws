package forms

import (
	"regexp"

	"github.com/gookit/validate"
)

type UserForm struct {
	Mobile     string `json:"mobile" validate:"required|MobileValidate"`
	Password   string `json:"password" validate:"required|minLen:3|maxLen:20|PwdCheck"`
	RePassword string `json:"re_password" validate:"required|minLen:3|maxLen:20|RepeatPwdValidate"`
	CaptchaId  string `json:"captcha_id" validate:"required"`
	Code       string `json:"code" validate:"required|len:6|ValidateCaptcha"`
}

// ConfigValidation 初始配置验证器
func (f UserForm) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"login":       []string{"Mobile", "Password"},
		"getUserInfo": []string{"Mobile"},
	})
}

// Messages 您可以自定义验证器错误消息
func (f UserForm) Messages() map[string]string {
	return validate.MS{
		"Mobile.MobileValidate":        "{field}输入不正确",
		"Code.ValidateCaptcha":         "{field}输入不正确",
		"RePassword.RepeatPwdValidate": "两次密码不一致",
		"Password.PwdCheck":            "密码中含有除-,_,&,@等特殊符号外的其他符号",
	}
}

// Translates 你可以自定义字段翻译
func (f UserForm) Translates() map[string]string {
	return validate.MS{
		"Mobile":     "手机号码",
		"Password":   "密码",
		"RePassword": "重复密码",
		"CaptchaId":  "验证码ID",
		"Code":       "验证码",
	}
}

func (f UserForm) MobileValidate(val string) bool {
	ok, _ := regexp.MatchString(`^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`, val)
	if !ok {
		return false
	}
	return true
}

func (f UserForm) RepeatPwdValidate(val string) bool {
	return f.Password == f.RePassword
}

func (f UserForm) PwdCheck(val string) bool {
	rxAlphaDash := regexp.MustCompile(`^(?:[\w-&@]+)$`)
	return rxAlphaDash.MatchString(val)
}

func (f UserForm) ValidateCaptcha(val string) bool {
	//return utils.Store.Verify(f.CaptchaId,f.Code,true)
	return true
}
