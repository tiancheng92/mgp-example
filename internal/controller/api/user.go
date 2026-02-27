package api

import (
	"mgp_example/internal/service"
	"mgp_example/internal/type/request"
	"net/http"

	"github.com/tiancheng92/mgp"
)

type userController struct {
	svc service.UserServiceInterface
}

func NewUserRouter(group *mgp.RouterGroup) {
	c := &userController{service.NewUserService()}
	g := group.Group("user").SetTagsForSwagger("用户")
	{
		g.GET(":auth_type", c.GetAuthList).
			SetSummaryForSwagger("获取用户权限列表").
			SetPathForSwagger(new(request.AuthType)).
			SetReturnsForSwagger(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[[]string]),
			})
		g.GET("", c.GetUserInfo).
			SetSummaryForSwagger("获取用户信息").
			SetReturnsForSwagger(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[string]),
			})
	}
}

func (c *userController) GetAuthList(ctx *mgp.Context) {
	p := new(request.AuthType)
	ctx.BindParams(p).HR(func() (any, error) {
		return c.svc.GetAuthList(ctx, p.AuthType)
	})
}

func (c *userController) GetUserInfo(ctx *mgp.Context) {
	ctx.HR(func() (any, error) {
		return c.svc.GetUserInfo(ctx)
	})
}
