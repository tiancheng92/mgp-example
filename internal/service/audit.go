package service

import (
	"fmt"
	"mgp_example/internal/service/audit"
	"mgp_example/internal/store/model"
	"mgp_example/internal/store/repository"
	"mgp_example/pkg/log"

	"github.com/bytedance/sonic"
	"github.com/tiancheng92/mgp"
)

type auditService struct {
	*readOnlyGenericService[model.Audit]
	userSvc   UserServiceInterface
	auditRepo repository.AuditRepoInterface
}

func NewAuditService() AuditServiceInterface {
	auditRepo := repository.NewAuditRepository()
	return &auditService{newReadOnlyGenericService[model.Audit](auditRepo), NewUserService(), auditRepo}
}

type AuditInfo struct {
	ID uint64
}

func (ai *AuditInfo) Handle(err error) {
	if ai.ID == 0 {
		return
	}

	var (
		status model.AuditStatus
		reason string
	)

	if err != nil {
		status = model.AuditFailed
		reason = err.Error()
	} else {
		status = model.AuditSuccess
	}

	go func() {
		if err = repository.NewAuditRepository().UpdateStatus(ai.ID, status, reason); err != nil {
			log.Errorf("update audit status failed: %v", err)
		}
	}()
}

func (s *auditService) New(ctx *mgp.Context, operation string, opts ...audit.Option) (*AuditInfo, error) {
	var setting audit.Setting
	for i := range opts {
		opts[i](&setting)
	}

	user, err := s.userSvc.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	odb, err := sonic.Marshal(setting.OriginalData)
	if err != nil {
		return nil, err
	}

	ndb, err := sonic.Marshal(setting.NewData)
	if err != nil {
		return nil, err
	}

	content, err := sonic.Marshal(setting.Content)
	if err != nil {
		return nil, err
	}

	data, err := s.auditRepo.Create(model.Audit{
		Username:     fmt.Sprintf("%s(%s)", user.Name, user.Username),
		Operation:    operation,
		Content:      content,
		OriginalData: odb,
		NewData:      ndb,
		Status:       model.AuditRunning,
	})
	if err != nil {
		return nil, err
	}

	ai := &AuditInfo{
		ID: data.ID,
	}

	ctx.Set("audit_info", ai)
	return ai, nil
}
