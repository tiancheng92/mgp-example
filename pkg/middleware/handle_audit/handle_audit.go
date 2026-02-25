package handle_audit

import (
	"mgp_example/service"

	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/mgp"
)

// HandleAudit 回写审计日志
func HandleAudit(ctx *gin.Context) {
	ctx.Next()
	if v, ok1 := ctx.Get("audit_info"); ok1 {
		if auditInfo, ok2 := v.(*service.AuditInfo); ok2 {
			var err error
			if allError, ok := ctx.Get(mgp.AllError); ok {
				if errData, ok3 := allError.(error); ok3 {
					err = errData
				}
			}
			go service.NewAuditService().Handle(auditInfo, err)
		}
	}
}
