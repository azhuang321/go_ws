package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/gin-gonic/gin"

	"ws_srv/errno"
)

// ErrReturn 错误快捷返回
func ErrReturn(c *gin.Context, errCode int, err *errno.Errno) {
	c.JSON(errCode, gin.H{
		"code": err.Code,
		"msg": err.Message,
		"data": []string{},
	})
}

// OkReturn 成功快捷返回
func OkReturn(c *gin.Context, data interface{},msg string) {
	successMsg := errno.OK.Message
	if msg != "" {
		successMsg = msg
	}
	c.JSON(http.StatusOK, gin.H{
		"code": errno.OK.Code,
		"msg": successMsg,
		"data": data,
	})
}


func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误" + e.Message(),
				})
			}
		}
	}
	return
}

func PrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("打印错误:%s", err.Error())
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Printf("打印错误:%s", err.Error())
		return
	}
	fmt.Println(out.String())
}