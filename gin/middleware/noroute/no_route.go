/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2022/10/17 15:31
@Description:
*/

package noroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twgcode/mbox/exception"
	"github.com/twgcode/mbox/gin/response"
)

func NoRoute(c *gin.Context) {
	response.Failed(c, exception.NewNotFound(""), response.WithHttpCode(http.StatusNotFound))
}
