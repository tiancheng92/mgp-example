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
	g := group.Group("audit").SetTags("Audit Record")
	{
		g.GET(":pk", c.Get).
			SetSummary("Get audit record by ID").
			SetPath(new(request.PrimaryKey)).
			SetReturns(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[model.Audit]),
			})
		g.GET("", c.List).
			SetSummary("List audit records").
			SetQuery(new(mgp.PaginateQuery)).
			SetReturns(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[mgp.ResultPaginateData[[]model.Audit]]),
			})
	}
}
