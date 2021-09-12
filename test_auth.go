/**
 * Auth :   liubo
 * Date :   2021/9/12 18:12
 * Comment:
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func authTest(router *gin.Engine) {

	var root = router.Group("/api/v1")

	// 分组

	// 用户相关
	var user = root.Group("/user")
	user.Use(MyAuth()) // 加入验证
	{
		user.GET("/shop", func(c *gin.Context) {
			c.JSON(http.StatusOK, "shop")
		})
	}

	// 登录
	root.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, "login succ")
	})

}

func MyAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var id = ctx.Query("userid")
		if len(id) > 0 {
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}


