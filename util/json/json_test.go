/*
*
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/7/12 10:2313
@Description:
*/
package json

import (
	"reflect"
	"testing"
)

type TestS struct {
	Name string
}

func TestMarshalObjectFormat(t *testing.T) {
	tests := map[string]struct {
		input          interface{}
		expectedResult []byte
		expectedError  error
	}{
		"test1": {
			input:          nil,
			expectedResult: []byte(emptyObjectFormatJSON),
			expectedError:  nil,
		},
		"test2": {
			input:          map[string]interface{}{"key1": "value1", "key2": "value2"},
			expectedResult: []byte(`{"key1":"value1","key2":"value2"}`),
			expectedError:  nil,
		},
		"test3": {
			input:          (*TestS)(nil),
			expectedResult: []byte(`{}`),
			expectedError:  nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := MarshalObjectFormat(test.input)

			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Errorf("For input %#v, expected result '%s', but got '%s'",
					test.input, test.expectedResult, result)
			}

			if err != test.expectedError {
				t.Errorf("For input %#v, expected error '%v', but got '%v'",
					test.input, test.expectedError, err)
			}
		})
	}
}
