package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"

	"user_api/errno"
	"user_api/utils"
)

// GetCaptcha 获取验证码
func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, utils.Store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误:%s", err.Error())
		rErr := errno.ErrRuntime.ReplaceMsg("生成验证码错误")
		utils.ErrReturn(c, http.StatusInternalServerError, &rErr)
		return
	}
	utils.OkReturn(c, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	}, "")
}
