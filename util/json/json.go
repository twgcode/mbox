/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/7/12 10:1004
@Description:
*/

package json

import (
	"encoding/json"
	"reflect"
)

const (
	emptyObjectFormatJSON = "{}"
)

// MarshalObjectFormat Object 格式
func MarshalObjectFormat(v interface{}) (data []byte, err error) {
	if v == nil {
		data = []byte(emptyObjectFormatJSON)
		return
	}
	if reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil() {
		return []byte(emptyObjectFormatJSON), nil
	}

	data, err = json.Marshal(v)
	return

}
