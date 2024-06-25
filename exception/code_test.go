/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/12/5 11:04
@Description:
*/

package exception

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

func TestCodeToHttpCode(t *testing.T) {
	tests := []struct {
		code     Code
		expected int
	}{
		{Success, http.StatusOK},
		{UnKnownException, http.StatusInternalServerError},
		{InvalidParam, http.StatusBadRequest},
		{IllegalOperation, http.StatusBadRequest},
		{AccessTokenExpired, http.StatusUnauthorized},
		{RefreshTokenExpired, http.StatusUnauthorized},
		{AccessTokenIllegal, http.StatusUnauthorized},
		{RefreshTokenIllegal, http.StatusUnauthorized},
		{NotFound, http.StatusNotFound},             // 假设 CustomErrorCode 在代码中未定义
		{Code(999), http.StatusInternalServerError}, // 假设 CustomErrorCode 在代码中未定义
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(int64(tt.code), 10), func(t *testing.T) {
			result := CodeToHttpCode(tt.code)
			if result != tt.expected {
				t.Errorf("期望为 %d，但对于代码 %d，得到了 %d", tt.expected, tt.code, result)
			}
		})
	}
}

func TestErrToHttpCode(t *testing.T) {
	tests := map[string]struct {
		err      error
		expected int
	}{
		"InternalServerError": {NewInternalServerError("测试"), http.StatusInternalServerError},
		"InvalidParam":        {NewInvalidParam("测试"), http.StatusBadRequest},
		"IllegalOperation":    {NewIllegalOperation("测试"), http.StatusBadRequest},
		"50000":               {NewAPIException(Code(50000), "测试"), http.StatusInternalServerError},
		"其他错误":                {fmt.Errorf("其他错误"), http.StatusInternalServerError},
		"nil":                 {nil, http.StatusOK},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := ErrToHttpCode(tt.err)
			if result != tt.expected {
				t.Errorf("期望为 %d, 得到了 %d", tt.expected, result)
			}
		})
	}
}
