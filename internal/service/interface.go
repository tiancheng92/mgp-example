package service

import (
	"mgp_example/internal/service/audit"
	"mgp_example/internal/store/model"

	ump_sdk "github.com/Yostardev/ump-sdk"
	"github.com/tiancheng92/mgp"
)

type GenericInterface[M model.Interface] interface {
	ReadOnlyGenericInterface[M]
	Update(ctx *mgp.Context, pk any, request any) (*M, error)
	Create(ctx *mgp.Context, request any) (*M, error)
	Delete(ctx *mgp.Context, pk any) error
}

type ReadOnlyGenericInterface[M model.Interface] interface {
	Get(ctx *mgp.Context, pk any) (*M, error)
	List(ctx *mgp.Context, pq *mgp.PaginateQuery) (*mgp.PaginateData[M], error)
	All(ctx *mgp.Context) ([]*M, error)
	Count(ctx *mgp.Context) (int64, error)
	Distinct(ctx *mgp.Context, field string) ([]string, error)
}

type AuditServiceInterface interface {
	ReadOnlyGenericInterface[model.Audit]
	New(ctx *mgp.Context, operation string, opts ...audit.Option) (*AuditInfo, error)
}

type UserServiceInterface interface {
	GetUserInfo(ctx *mgp.Context) (*ump_sdk.UserInfo, error)
	GetAuthList(ctx *mgp.Context, authType string) ([]string, error)
}
