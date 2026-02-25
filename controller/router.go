package controller

import (
	"mgp_example/controller/api"
	"mgp_example/controller/api/universal"
	_ "mgp_example/docs"
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

	r.GET("healthz", universal.HealthCheck)

	r.GET("/swagger/*any", ginSwagger.WrapHandler())

	apiGroup := r.Group("/api", auth.Check)
	{
		api.NewUserRouter(apiGroup)
		api.NewAuditRouter(apiGroup)
	}
	return r
}
