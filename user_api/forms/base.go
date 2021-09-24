package forms

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"

	"user_api/errno"
	"user_api/utils"
)

func validateErr(errStr string) *errno.Errno {
	validateErr := *errno.ErrRequestParams
	validateErr.Message = errStr
	return &validateErr
}

// BindJsonAndValidate 绑定参数并验证
func BindJsonAndValidate(c *gin.Context, obj interface{},scene string) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		rErr := errno.ErrBind.ReplaceMsg("参数错误")
		utils.ErrReturn(c, http.StatusBadRequest, &rErr)
		return false
	}
	v := validate.Struct(obj,scene)
	if !v.Validate() {
		utils.ErrReturn(c, http.StatusBadRequest, validateErr(v.Errors.One()))
		return false
	}
	return true
}
