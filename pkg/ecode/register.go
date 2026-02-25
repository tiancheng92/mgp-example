package ecode

import (
	"github.com/tiancheng92/mgp/errors"
)

// 客户端错误
const (
	ErrClientAuth         = "ErrClientAuth"
	ErrClientPageNotFound = "ErrClientPageNotFound"
	ErrClientDataNotFound = "ErrClientDataNotFound"
)

// 服务端错误
const (
	ErrServerGet          = "ErrServerGet"
	ErrServerCreate       = "ErrServerCreate"
	ErrServerUpdate       = "ErrServerUpdate"
	ErrServerDelete       = "ErrServerDelete"
	ErrServerDuplicateKey = "ErrServerDuplicateKey"
)

func init() {
	errors.Register(ErrClientAuth, 403001, 403, "权限异常")
	errors.Register(ErrClientPageNotFound, 404001, 404, "页面/接口不存在")
	errors.Register(ErrClientDataNotFound, 404002, 404, "数据不存在")
	errors.Register(ErrServerGet, 500001, 500, "数据获取失败")
	errors.Register(ErrServerCreate, 500002, 500, "创建失败")
	errors.Register(ErrServerUpdate, 500003, 500, "更新失败")
	errors.Register(ErrServerDelete, 500004, 500, "删除失败")
	errors.Register(ErrServerDuplicateKey, 500005, 500, "数据已存在或与现有数据冲突")
}
