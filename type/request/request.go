package request

import "mgp_example/store/model"

type Interface interface {
	FormatToModel() model.Interface
}

type PrimaryKey struct {
	PrimaryKey uint64 `uri:"pk" binding:"required"`
}

type DistinctField struct {
	Field string `uri:"field" binding:"required"`
}

type Header struct {
	Authorization string `form:"token" header:"Authorization" binding:"required"`
}

type AuthType struct {
	AuthType string `uri:"auth_type" binding:"required"`
}
