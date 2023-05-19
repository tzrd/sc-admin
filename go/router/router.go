package router

import (
	md "scadmin/middleware"
	routerV1 "scadmin/router/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	v1 := e.Group("/api/v1/")
	// 鉴权路由
	v1.Use(md.JWTAuth())

	// sys
	routerV1.LoadSys(v1)
	routerV1.LoadMfa(v1)
}
