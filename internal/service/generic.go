package service

import (
	"mgp_example/internal/service/audit"
	"mgp_example/internal/store/model"
	"mgp_example/internal/store/repository"
	"mgp_example/internal/type/request"
	"reflect"

	"github.com/Yostardev/gf"
	"github.com/tiancheng92/mgp"
)

type genericService[M model.Interface] struct {
	*readOnlyGenericService[M]
	repo  repository.GenericInterface[M]
	audit AuditServiceInterface
}

func newGenericService[M model.Interface](repo repository.GenericInterface[M]) *genericService[M] {
	return &genericService[M]{
		newReadOnlyGenericService[M](repo),
		repo,
		NewAuditService(),
	}
}

func toModel[M model.Interface](requestPtr any) M {
	switch d := requestPtr.(type) {
	case request.Interface:
		return d.FormatToModel().(M)
	default:
		modelPtr := new(M)
		r := reflect.ValueOf(requestPtr).Elem()
		m := reflect.ValueOf(modelPtr).Elem()

		if r.Kind() != reflect.Struct || m.Kind() != reflect.Struct {
			panic("Both model and request input must be structs ptr")
		}

		for i := range r.NumField() {
			requestField := r.Type().Field(i)
			modelFieldValue := m.FieldByName(requestField.Name)
			if modelFieldValue.IsValid() && requestField.Type == modelFieldValue.Type() {
				modelFieldValue.Set(r.Field(i))
			}
		}
		return *modelPtr
	}
}

func (s *genericService[M]) Create(ctx *mgp.Context, request any) (*M, error) {
	if err := s.audit.New(
		ctx,
		gf.StringJoin("Create ", (*new(M)).GetTableName()),
		audit.WithNewData(request),
	); err != nil {
		return nil, err
	}

	return s.repo.Create(toModel[M](request))
}

func (s *genericService[M]) Update(ctx *mgp.Context, pk any, request any) (*M, error) {
	originalData, err := s.repo.Get(pk)
	if err != nil {
		return nil, err
	}

	if err = s.audit.New(
		ctx,
		gf.StringJoin("Update ", (*new(M)).GetTableName()),
		audit.WithContent(audit.CM{"id": pk}),
		audit.WithData(originalData, request),
	); err != nil {
		return nil, err
	}
	return s.repo.Update(pk, toModel[M](request))
}

func (s *genericService[M]) Delete(ctx *mgp.Context, pk any) error {
	originalData, err := s.repo.Get(pk)
	if err != nil {
		return err
	}

	if err = s.audit.New(
		ctx,
		gf.StringJoin("Delete ", (*new(M)).GetTableName()),
		audit.WithContent(audit.CM{"id": pk}),
		audit.WithOriginalData(originalData),
	); err != nil {
		return err
	}

	return s.repo.Delete(pk)
}

// ------------------------------

type readOnlyGenericService[M model.Interface] struct {
	repo repository.GenericInterface[M]
}

func newReadOnlyGenericService[M model.Interface](repo repository.GenericInterface[M]) *readOnlyGenericService[M] {
	return &readOnlyGenericService[M]{repo}
}

func (ros *readOnlyGenericService[M]) Get(_ *mgp.Context, pk any) (*M, error) {
	return ros.repo.Get(pk)
}

func (ros *readOnlyGenericService[M]) List(_ *mgp.Context, pq *mgp.PaginateQuery) (*mgp.PaginateData[M], error) {
	return ros.repo.List(pq)
}

func (ros *readOnlyGenericService[M]) All(ctx *mgp.Context) ([]*M, error) {
	return ros.repo.All(ctx.Request.URL.Query())
}

func (ros *readOnlyGenericService[M]) Count(ctx *mgp.Context) (int64, error) {
	return ros.repo.Count(ctx.Request.URL.Query())
}

func (ros *readOnlyGenericService[M]) Distinct(_ *mgp.Context, field string) ([]string, error) {
	return ros.repo.Distinct(field)
}
