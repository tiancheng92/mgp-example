package handle_error

import (
	"mgp_example/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/mgp"
)

// HandleError 错误处理
func HandleError(ctx *gin.Context) {
	ctx.Next()
	if warnInfo, exist := ctx.Get(mgp.ErrorLogLevelWarn); exist {
		log.Warnf("%v", warnInfo)
	}
	if errInfo, exist := ctx.Get(mgp.ErrorLogLevelError); exist {
		log.Errorf("%+v", errInfo)
	}
}
