package model

import (
	"time"

	"gorm.io/gorm"
)

type Interface interface {
	GetPrimaryKeyName() string
	GetFuzzySearchFieldList() []string
	GetDefaultOrder() string
	GetTableName() string
}

type Model struct {
	ID        uint64    `json:"id" gorm:"primary_key;type:bigint unsigned;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (Model) GetPrimaryKeyName() string {
	return "id"
}

func (Model) GetFuzzySearchFieldList() []string {
	return []string{}
}

func (Model) GetDefaultOrder() string {
	return "id desc"
}

type SoftDeleteModel struct {
	Model
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
