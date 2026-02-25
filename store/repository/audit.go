package repository

import (
	"mgp_example/pkg/ecode"
	"mgp_example/store/model"

	"github.com/tiancheng92/mgp/errors"
)

type auditRepository struct {
	*genericRepository[model.Audit]
}

func NewAuditRepository() AuditRepoInterface {
	return &auditRepository{newGenericRepository[model.Audit]()}
}

func (r *auditRepository) UpdateStatus(id uint64, status model.AuditStatus, failedReason string) error {
	err := r.db.Where("`id` = ?", id).Select("status", "failed_reason").Updates(&model.Audit{
		Status:       status,
		FailedReason: failedReason,
	}).Error
	return errors.WithCode(ecode.ErrServerUpdate, err)
}
