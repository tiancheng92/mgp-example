package repository

import (
	"mgp_example/internal/store/model"
	"net/url"

	"github.com/tiancheng92/mgp"
)

type GenericInterface[M model.Interface] interface {
	Get(pk any) (*M, error)
	Create(attributes M) (*M, error)
	Update(pk any, attributes M) (*M, error)
	Delete(pk any) error
	List(pq *mgp.PaginateQuery) (*mgp.PaginateData[M], error)
	Count(params url.Values) (int64, error)
	Distinct(field string) ([]string, error)
	All(params url.Values) ([]*M, error)
}

type AuditRepoInterface interface {
	GenericInterface[model.Audit]
	UpdateStatus(id uint64, status model.AuditStatus, reason string) error
}
