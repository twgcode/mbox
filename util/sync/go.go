/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/4/26 13:44
@Description:
*/

package sync

import (
	"go.uber.org/zap"
)

// Go 开启 goroutine
func Go(f func(), desc string, l *zap.Logger) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if l == nil {
					l = zap.L()
				}
				l.Error("goroutine panic", zap.String("desc", desc), zap.Any("error", err))
			}
		}()
		f()
	}()
}
