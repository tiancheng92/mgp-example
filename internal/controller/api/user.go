package api

import (
	"mgp_example/internal/service"
	"mgp_example/internal/type/request"

	"github.com/tiancheng92/mgp"
)

type userController struct {
	svc service.UserServiceInterface
}

func NewUserRouter(group *mgp.RouterGroup) {
	c := &userController{service.NewUserService()}
	g := group.Group("user").SwaggerTags("用户")
	{
		g.GET(":auth_type", c.GetAuthList).
			SwaggerSummary("获取用户权限列表").
			SwaggerPath(new(request.AuthType)).
			SwaggerReturns(mgp.RT[[]string]())
		g.GET("", c.GetUserInfo).
			SwaggerSummary("获取用户信息").
			SwaggerReturns(mgp.RT[string]())
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
