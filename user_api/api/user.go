package api

import (
	"github.com/gin-gonic/gin"
	"user_api/forms"

	"user_api/forms/user"
)

func Register(ctx *gin.Context) {
	registerForm := user.RegisterFrom{}
	if ok := forms.BindJsonAndValidate(ctx, &registerForm); !ok {
		return
	}


}

func Login(ctx *gin.Context) {

}


