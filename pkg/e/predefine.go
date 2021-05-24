package e

import "errors"

var (
	ErrNotFound = errors.New("数据查询为空")
	ErrInternal = errors.New("内部错误")
)
