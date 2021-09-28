package forms

import (
	"github.com/gookit/validate"
)

type UserForm struct {
	UserId string `json:"mobile" validate:"required|MobileValidate"`
}

// ConfigValidation 初始配置验证器
func (f UserForm) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"login": []string{"Mobile", "Password"},
	})
}

// Messages 您可以自定义验证器错误消息
func (f UserForm) Messages() map[string]string {
	return validate.MS{
		"Mobile.MobileValidate": "{field}输入不正确",
	}
}

// Translates 你可以自定义字段翻译
func (f UserForm) Translates() map[string]string {
	return validate.MS{
		"Mobile": "手机号码",
	}
}
