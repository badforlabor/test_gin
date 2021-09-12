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
	globalRouter = Routers()

	authTest(globalRouter)
	cookieTest(globalRouter)

	globalRouter.Run(*addr)
}

func Routers() *gin.Engine {
	var router = gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}
