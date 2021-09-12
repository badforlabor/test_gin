/**
 * Auth :   liubo
 * Date :   2021/9/12 18:16
 * Comment:
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func cookieTest(router *gin.Engine) {

	var root = router.Group("/api/v1")

	// 分组

	// 用户相关
	var user = root.Group("/cookie")
	{
		user.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, "cookie info")
		})
	}

}



