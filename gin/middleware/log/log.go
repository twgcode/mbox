/*
@Author: twgcode
@Email:  17600113577@163.com
@Date:   2022/10/16 10:30 PM
@Description:
*/

package log

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 接收gin框架默认的日志
func GinLogger(mode string, lg *zap.Logger) gin.HandlerFunc {
	// debug 模式 只控制台只输出一份
	if mode == gin.DebugMode {
		return func(c *gin.Context) {
			path := c.Request.URL.Path
			query := c.Request.URL.RawQuery
			lg.Info(path,
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			)
			c.Next()
		}
	}
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		lg.Info(path,
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
		c.Next()
		cost := time.Since(start)
		lg.Info(path,
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("cost", cost),
		)
	}
}
