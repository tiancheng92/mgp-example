package controller

import (
	_ "mgp_example/docs"
	"mgp_example/internal/controller/api"
	"mgp_example/internal/controller/api/universal"
	"mgp_example/pkg/log"
	"mgp_example/pkg/middleware/auth"
	"mgp_example/pkg/middleware/cross_domain"
	"mgp_example/pkg/middleware/handle_audit"
	"mgp_example/pkg/middleware/handle_error"
	_ "mgp_example/pkg/validator"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/tiancheng92/gin-swagger"
	"github.com/tiancheng92/mgp"
)

// InitRouter 初始化路由
func InitRouter() *mgp.Engine {
	mgp.SetSuccessMsg("OK")
	mgp.SetSuccessCode(0)

	r := mgp.New()

	r.Use(
		ginzap.GinzapWithConfig(log.GetLogger(), &ginzap.Config{
			TimeFormat: time.DateTime,
			UTC:        true,
			SkipPaths: []string{
				"/healthz",
			},
		}),
		gin.Recovery(),
		handle_audit.HandleAudit,
		handle_error.HandleError,
		cross_domain.CrossDomain(),
	)

	r.NoRoute(universal.NoRoute)

	r.NoMethod(universal.NoMethod)

	r.RawGET("/healthz", universal.HealthCheck)

	r.RawGET("/swagger/*any", ginSwagger.WrapHandler())

	r.GET("", func(c *mgp.Context) {
		c.HR(func() (any, error) {
			return 1, nil
		})
	}).SetTags("tmp").SetReturns(&mgp.ReturnType{
		StatusCode: 200,
		Body:       new(mgp.Result[int]),
	})

	apiGroup := r.Group("/api", auth.Check).SetUseApiKeyAuth()
	{
		api.NewUserRouter(apiGroup)
		api.NewAuditRouter(apiGroup)
	}
	return r
}
