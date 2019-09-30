package tsValid

import (
	"fmt"
	"reflect"
)

// 验证错误提示
type ValidateError struct {
	Field string // 验证不通过的属性
	Valid string // 不通过的条件
}

func (this ValidateError) Error() string {
	return fmt.Sprintf("field:%s Unable to verify:%s", this.Field, this.Valid)
}

// 校验封装
type ValidateInterface interface {
	validate() bool
}

type validateModel struct {
	required bool
	fieldT   reflect.StructField
	fieldV   interface{}
}
