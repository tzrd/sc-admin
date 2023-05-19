package router

import (
	"scadmin/controller"

	"github.com/gin-gonic/gin"
)

func WhiteRouter(e *gin.Engine) {
	r := e.Group("/api/v1/other/")
	{
		r.POST("/login", controller.UserLogin)
	}
}
