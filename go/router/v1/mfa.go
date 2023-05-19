package routerV1

import (
	ControllerMfa "scadmin/controller/mfa"

	"github.com/gin-gonic/gin"
)

func LoadMfa(e *gin.RouterGroup) {
	{
		// MFA
		//e.GET("/mfa", ControllerMfa.MfaList)
		e.GET("/mfa", ControllerMfa.MfaSearch)
		e.PUT("/mfa/:id", ControllerMfa.MfaUpdate)
		e.POST("/mfa", ControllerMfa.MfaCreate)
		e.DELETE("/mfa/:id", ControllerMfa.MfaDel)
		e.GET("/mfa/getInfo", ControllerMfa.GetMfaInfo)

		// OTP
		e.GET("/otp", ControllerMfa.OtpSearch)
	}
}
