/**
 * Auth :   liubo
 * Date :   2021/9/12 17:52
 * Comment:
 */

package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var globalRouter *gin.Engine

var addr = flag.String("addr", ":8080", ":8080")

func main() {
	flag.Parse()
	initLog()

	initGin()

	authTest(globalRouter)
	cookieTest(globalRouter)

	globalRouter.Run(*addr)
}

func initGin() {
	var useDefault = true
	var router *gin.Engine
	if useDefault {
		router = gin.Default()
	} else {
		router = gin.New()
		// router.Use(gin.Logger())
		useLogMiddle(globalRouter)
		router.Use(gin.Recovery())
	}

	globalRouter = router
}

