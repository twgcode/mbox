/*
@Author: twgcode
@Email:  17600113577@163.com
@Date:   2022/10/16 10:40 PM
@Description:
*/

package redisync

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/twgcode/tools/sync"

	"github.com/twgcode/mbox/exception"
	"github.com/twgcode/mbox/gin/response"
)

// NoWaitMiddleware  非等待/阻塞 的锁
func NoWaitMiddleware(client *redis.ClusterClient, key string, expiration time.Duration) (f func(c *gin.Context), err error) {
	var (
		locker *sync.NoWaitLock
	)
	if locker, err = sync.NewNoWaitLock(client, key, expiration); err != nil {
		return
	}
	f = func(c *gin.Context) {
		var (
			err1 error
		)
		_, err1 = locker.Do(func() error {
			c.Next()
			return nil
		})
		if err1 != nil { // 错误处理
			if errors.Is(err1, sync.ErrLockedState) {
				response.Failed(c, exception.NewConflict("the current lock is locked"))
			} else {
				response.Failed(c, exception.NewInternalServerError(err1.Error()), response.WithInternalServerError())
			}
			c.AbortWithStatus(http.StatusOK)
			return
		}
	}
	return
}
