package api

import (
	"mgp_example/service"
	"mgp_example/store/model"
	"mgp_example/type/request"

	"github.com/tiancheng92/mgp"
)

type genericController[R any, M model.Interface] struct {
	*readOnlyGenericController[M]
	service.GenericInterface[M]
}

func newGenericController[R any, M model.Interface](service service.GenericInterface[M]) *genericController[R, M] {
	return &genericController[R, M]{newReadOnlyGenericController[M](service), service}
}

func (c *genericController[R, M]) Get(ctx *mgp.Context) {
	c.readOnlyGenericController.Get(ctx)
}

func (c *genericController[R, M]) List(ctx *mgp.Context) {
	c.readOnlyGenericController.List(ctx)
}

func (c *genericController[R, M]) All(ctx *mgp.Context) {
	c.readOnlyGenericController.All(ctx)
}

func (c *genericController[R, M]) Create(ctx *mgp.Context) {
	r := new(R)
	ctx.BindBody(r).HR(func() (any, error) {
		return c.GenericInterface.Create(ctx, r)
	})
}

func (c *genericController[R, M]) Update(ctx *mgp.Context) {
	p := new(request.PrimaryKey)
	r := new(R)
	ctx.BindParams(p).BindBody(r).HR(func() (any, error) {
		return c.GenericInterface.Update(ctx, p.PrimaryKey, r)
	})
}

func (c *genericController[R, M]) Delete(ctx *mgp.Context) {
	p := new(request.PrimaryKey)
	ctx.BindParams(p).HR(func() error {
		return c.GenericInterface.Delete(ctx, p.PrimaryKey)
	})
}

// ------------------------------

type readOnlyGenericController[M model.Interface] struct {
	service.ReadOnlyGenericInterface[M]
}

func newReadOnlyGenericController[M model.Interface](service service.ReadOnlyGenericInterface[M]) *readOnlyGenericController[M] {
	return &readOnlyGenericController[M]{service}
}

func (roc *readOnlyGenericController[M]) Get(ctx *mgp.Context) {
	p := new(request.PrimaryKey)
	ctx.BindParams(p).HR(func() (any, error) {
		return roc.ReadOnlyGenericInterface.Get(ctx, p.PrimaryKey)
	})
}

func (roc *readOnlyGenericController[M]) List(ctx *mgp.Context) {
	pq := new(mgp.PaginateQuery)
	ctx.BindPaginateQuery(pq).HR(func() (any, error) {
		return roc.ReadOnlyGenericInterface.List(ctx, pq)
	})
}

func (roc *readOnlyGenericController[M]) All(ctx *mgp.Context) {
	ctx.HR(func() (any, error) {
		return roc.ReadOnlyGenericInterface.All(ctx)
	})
}

func (roc *readOnlyGenericController[M]) Count(ctx *mgp.Context) {
	ctx.HR(func() (any, error) {
		return roc.ReadOnlyGenericInterface.Count(ctx)
	})
}

func (roc *readOnlyGenericController[M]) Distinct(ctx *mgp.Context) {
	p := new(request.DistinctField)
	ctx.BindParams(p).HR(func() (any, error) {
		return roc.ReadOnlyGenericInterface.Distinct(ctx, p.Field)
	})
}
