/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/4/12 12:15
@Description:
*/

package exception

import (
	"testing"
)

func TestNewAPIException(t *testing.T) {
	tests := []struct {
		name    string
		code    Code
		format  string
		a       []interface{}
		errFunc func(format string, a ...interface{}) APIException
	}{
		{name: "Unauthorized", code: Unauthorized, format: "无权限: %s, %q", a: []interface{}{"key", "xx"}, errFunc: NewUnauthorized},
		{name: "BadRequest", code: BadRequest, format: "无权限: %s, %q", a: []interface{}{"key", "xx"}, errFunc: NewBadRequest},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			err1 := NewAPIException(tt.code, tt.format, tt.a...)
			err2 := tt.errFunc(tt.format, tt.a...)
			if err1.ErrorCode() != err2.ErrorCode() || err1.Error() != err2.Error() {
				t.Errorf("err1: %q, code: %q, err2: %q, code: %q", err1, CodeReason(err1.ErrorCode()),
					err2, CodeReason(err2.ErrorCode()))
			}
		})
	}
}
