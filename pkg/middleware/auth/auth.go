package auth

import (
	"mgp_example/config"
	"mgp_example/internal/type/request"
	"mgp_example/pkg/ecode"

	ump "github.com/Yostardev/ump-sdk"
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/mgp"
	"github.com/tiancheng92/mgp/errors"
)

func Check(ctx *gin.Context) {
	header := new(request.Header)
	if err := ctx.ShouldBindHeader(header); err != nil {
		mgp.Response(ctx, nil, mgp.HandleValidationErr(err))
		return
	}

	ctx.Set("token", header.Authorization)

	if ok, err := ump.NewClient(config.GetConf().Ump.Url, config.GetConf().Ump.APIAppID, header.Authorization).
		CheckAuth(ctx.Request.URL.RequestURI(), ctx.Request.Method); err != nil || !ok {
		mgp.Response(ctx, nil, errors.WithCode(ecode.ErrClientAuth, "Error Auth Check"))
		return
	}

	ctx.Next()
}

func CheckWebsocket(ctx *gin.Context) {
	token, err := ctx.Cookie(config.GetConf().Ump.CookieName)
	if err != nil {
		mgp.Response(ctx, nil, mgp.HandleValidationErr(err))
		return
	}

	ctx.Set("token", token)

	if ok, err := ump.NewClient(config.GetConf().Ump.Url, config.GetConf().Ump.APIAppID, token).
		CheckAuth(ctx.Request.URL.RequestURI(), ctx.Request.Method); err != nil || !ok {
		mgp.Response(ctx, nil, errors.WithCode(ecode.ErrClientAuth, "Error Auth Check"))
		return
	}

	ctx.Next()
}
