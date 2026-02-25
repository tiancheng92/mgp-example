package api

import (
	"mgp_example/service"
	"mgp_example/store/model"
	"mgp_example/type/request"
	"net/http"

	"github.com/tiancheng92/mgp"
)

type auditController struct {
	*readOnlyGenericController[model.Audit]
}

func NewAuditRouter(group *mgp.RouterGroup) {
	c := &auditController{newReadOnlyGenericController[model.Audit](service.NewAuditService())}
	g := group.Group("audit")
	{
		g.GET(":pk", c.Get).
			SetTags("Audit Record").
			SetSummary("Get audit record by ID").
			SetPath(new(request.PrimaryKey)).
			SetReturns(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[model.Audit]),
			})
		g.GET("", c.List).
			SetTags("Audit Record").
			SetSummary("List audit records").
			SetQuery(new(mgp.PaginateQuery)).
			SetReturns(&mgp.ReturnType{
				StatusCode: http.StatusOK,
				Body:       new(mgp.Result[mgp.ResultPaginateData[[]model.Audit]]),
			})
	}
}
