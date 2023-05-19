package ControllerMfa

import (
	ModelMfa "scadmin/model/mfa"
	ServiceMfa "scadmin/service/mfa"
	"scadmin/utils"

	"github.com/gin-gonic/gin"
)

// @Tags OTP
// @Description 获取OTP列表
// @Param   page query	string false "page"
// @Param   size_size query	string false "page_size"
// @Success 200 {string} string	"ok"
// @Router /api/v1/otp [get]
func OtpList(c *gin.Context) {
	var list []ModelMfa.OtpInfo
	dates, result := ServiceMfa.OtpQuery(list, c)
	utils.R(dates, result, c)
}

// @Tags OTP
// @Description 获取OTP列表
// @Param   query	string true "q"
// @Param   page query	string false "page"
// @Param   size_size query	string false "page_size"
// @Success 200 {string} string	"ok"
// @Router /api/v1/otp [get]
func OtpSearch(c *gin.Context) {
	var list []ModelMfa.OtpInfo
	dates, result := ServiceMfa.OtpSearch(list, c)
	utils.R(dates, result, c)
}

// @Tags OTP
// @Description 获取单条OTP信息
// @Param   id path	string true "id"
// @Success 200 {string} string	"ok"
// @Router /api/v1/otp/{id} [get]
func OtpSinge(c *gin.Context) {
	id := c.Param("id")
	singData, result := ServiceMfa.OtpSinge(id)
	utils.R(singData, result, c)
}
