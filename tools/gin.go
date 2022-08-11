package tools

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type GinResponse struct {
	Result int         `json:"result"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error,omitempty"`
}

func (res *GinResponse) ErrorResponse(c *gin.Context, err error) {
	res.Result = 0
	res.Error = err.Error()

	fmt.Println(err.Error())
	if _, ok := errorList[err.Error()]; !ok {
		debug.PrintStack()
	}

	c.AbortWithStatusJSON(400, res)
}
