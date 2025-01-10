/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/4/26 13:44
@Description:
*/

package sync

import (
	"runtime/debug"

	"go.uber.org/zap"
)

// Go 开启 goroutine
func Go(f func(), desc string, l *zap.Logger, stack ...bool) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if l == nil {
					l = zap.L()
				}
				if len(stack) > 0 && stack[0] {
					l.Error("goroutine panic", zap.String("desc", desc), zap.Any("error", err), zap.Any("stack", string(debug.Stack())))
				} else {
					l.Error("goroutine panic", zap.String("desc", desc), zap.Any("error", err))
				}
			}
		}()
		f()
	}()
}
