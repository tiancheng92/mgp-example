package api

import (
	"mgp_example/internal/service"
	"mgp_example/internal/store/model"
	"mgp_example/internal/type/request"
	"net/http"

	"github.com/tiancheng92/mgp"
)

type auditController struct {
	*readOnlyGenericController[model.Audit]
}

func NewAuditRouter(group *mgp.RouterGroup) {
	c := &auditController{newReadOnlyGenericController[model.Audit](service.NewAuditService())}
	g := group.Group("audit").SetTagsForSwagger("审计记录")
	{
		g.GET(":pk", c.Get).
			SetSummaryForSwagger("获取指定ID的审计记录").
			SetPathForSwagger(new(request.PrimaryKey)).
			SetReturnsForSwagger(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[model.Audit]),
			})
		g.GET("", c.List).
			SetSummaryForSwagger("分页获取审计记录").
			SetQueryForSwagger(new(mgp.PaginateQuery)).
			SetReturnsForSwagger(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[mgp.ResultPaginateData[[]model.Audit]]),
			})
	}
}
