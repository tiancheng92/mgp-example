package api

import (
	"mgp_example/internal/service"
	"mgp_example/internal/store/model"
	"mgp_example/internal/type/request"

	"github.com/tiancheng92/mgp"
)

type auditController struct {
	*readOnlyGenericController[model.Audit]
}

func NewAuditRouter(group *mgp.RouterGroup) {
	c := &auditController{newReadOnlyGenericController[model.Audit](service.NewAuditService())}
	g := group.Group("audit").SwaggerTags("审计记录")
	{
		g.GET(":pk", c.Get).
			SwaggerSummary("获取指定ID的审计记录").
			SwaggerPath(new(request.PrimaryKey)).
			SwaggerReturns(mgp.RT[model.Audit]())
		g.GET("", c.List).
			SwaggerSummary("分页获取审计记录").
			SwaggerQuery(new(mgp.PaginateQuery)).
			SwaggerReturns(mgp.PRT[model.Audit]())
	}
}
