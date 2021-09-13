/**
 * Auth :   liubo
 * Date :   2021/9/12 18:16
 * Comment:
 */

package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func cookieTest(router *gin.Engine) {

	// cookie的秘钥
	store := cookie.NewStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	var root = router.Group("/api/v1")

	// 用户相关
	var user = root.Group("/cookie")
	{
		user.GET("/info", func(c *gin.Context) {
			var cookie, e = c.Request.Cookie("test-cookie")
			if e == nil && cookie != nil {
				// 此时，只有cookie.value，其他数值是没有的，比如cookie.MaxAge
			} else {
				if cookie != nil {
					mylog("expire time:", cookie.RawExpires)
				}
				cookie = &http.Cookie{Name:"test-cookie", Value:"0", MaxAge:3}
				http.SetCookie(c.Writer, cookie)
			}
			var v, _ = strconv.ParseInt(cookie.Value, 10, 0)
			v++
			cookie.Value = strconv.Itoa(int(v))
			cookie.MaxAge = 3	// 必须每次都设置因为收到客户端反馈后，并没有此数值，如果不设置，就是cookie不过期
			http.SetCookie(c.Writer, cookie)

			c.JSON(http.StatusOK, "cookie value=" + cookie.Value)
		})


		user.Use(sessions.Sessions("mysession", store))
		user.GET("/info2", func(c *gin.Context) {
			session := sessions.Default(c)
			var count int
			v := session.Get("count")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count++
			}
			session.Options(sessions.Options{MaxAge:3})
			session.Set("count", count)
			session.Save()

			c.JSON(http.StatusOK, "cookie value=" + strconv.Itoa(count))
		})

		// 跟info2类似。只不过使用了不同的key。但仍然使用的是同一个session。
		user.GET("/info3", func(c *gin.Context) {
			session := sessions.Default(c)
			var count int
			v := session.Get("v3-count")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count++
			}
			session.Options(sessions.Options{MaxAge:3})
			session.Set("v3-count", count)
			session.Save()

			c.JSON(http.StatusOK, "cookie value=" + strconv.Itoa(count))
		})
	}

}



