package universal

import (
	"mgp_example/pkg/ecode"

	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/mgp"
	"github.com/tiancheng92/mgp/errors"
)

// NoRoute 不匹配的路由返回错误
func NoRoute(ctx *gin.Context) {
	mgp.Response(ctx, nil, errors.WithCode(ecode.ErrClientPageNotFound, "Page not found"))
}

func NoMethod(ctx *gin.Context) {
	mgp.Response(ctx, nil, errors.WithCode(ecode.ErrClientMethNotAllow, "Method not allow"))
}
