/**
@Author: wei-g
@Date:   2021/6/23 8:18 下午
@Description:
*/

package crypto

import (
	"testing"
)

func TestHmacSha256(t *testing.T) {
	type test struct { // 定义test结构体
		key  string
		data string
		want string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"Short key": {key: "a", data: "test", want: "27524d31b31dda3501981299e2dea8266d63df43bdceb9772b49ea139d86efb1"},
		"Long key":  {key: "abcdef123456", data: "test", want: "e7fbd1f41b90a5dae1d22632c919fd7b71f3749e65bba58c07a6b476cb546fbf"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := HmacSha256(tc.key, tc.data)
			if tc.want != got {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
