package model

import (
	"github.com/tiancheng92/datatypes"
)

type Audit struct {
	Model
	Username     string         `json:"username" gorm:"type:varchar(64);not null;comment:用户"`
	Operation    string         `json:"operation" gorm:"type:varchar(64);not null;default:'';comment:操作"`
	Content      datatypes.JSON `json:"content" gorm:"type:json;not null;comment:上下文内容"`
	OriginalData datatypes.JSON `json:"original_data" gorm:"type:json;not null;comment:原始数据"`
	NewData      datatypes.JSON `json:"new_data" gorm:"type:json;not null;comment:新数据"`
	FailedReason string         `json:"failed_reason" gorm:"type:longtext;not null;comment:失败原因"`
	Status       AuditStatus    `json:"status" gorm:"type:enum('running', 'failed', 'success');not null;default:'running';comment:状态"`
}

func (Audit) GetFuzzySearchFieldList() []string {
	return []string{"username", "operation", "content", "status"}
}

func (Audit) GetTableName() string {
	return "Audit"
}

type AuditStatus string

const (
	AuditSuccess AuditStatus = "success"
	AuditFailed  AuditStatus = "failed"
	AuditRunning AuditStatus = "running"
)
