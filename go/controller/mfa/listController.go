package ControllerMfa

import (
	"net/http"
	"scadmin/configure"
	"scadmin/middleware"
	ModelMfa "scadmin/model/mfa"
	ServiceMfa "scadmin/service/mfa"
	"scadmin/utils"

	"github.com/gin-gonic/gin"
)

// @Tags MFA
// @Description 获取MFA列表
// @Param   page query	string false "page"
// @Param   size_size query	string false "page_size"
// @Success 200 {string} string	"ok"
// @Router /api/v1/mfa [get]
func MfaList(c *gin.Context) {
	var list []ModelMfa.MfaInfo
	dates, result := ServiceMfa.MfaQuery(list, c)
	utils.R(dates, result, c)
}

// @Tags MFA
// @Description 获取符合条件的MFA信息
// @Param   query	string true "q"
// @Param   page query	string false "page"
// @Param   size_size query	string false "page_size"
// @Success 200 {string} string	"ok"
// @Router /api/v1/mfa [get]
func MfaSearch(c *gin.Context) {
	var list []ModelMfa.MfaInfo
	dates, result := ServiceMfa.MfaSearch(list, c)
	utils.R(dates, result, c)
}

// @Tags MFA
// @Description 获取单条MFA信息
// @Param   id path	string true "id"
// @Success 200 {string} string	"ok"
// @Router /api/v1/mfa/{id} [get]
func MfaSinge(c *gin.Context) {
	id := c.Param("id")
	singData, result := ServiceMfa.MfaSinge(id)
	utils.R(singData, result, c)
}

// @Tags MFA
// @Description 添加MFA
// @Param account body string true "账号"
// @Param name body string true "名称"
// @Param desc body string false "描述"
// @Param status body int false "所属状态是否有效  1是有效 0是失效"
// @Success 200 {string} string	"ok"
// @Router /api/v1/mfa [post]
func MfaCreate(c *gin.Context) {
	var mfaInfo ModelMfa.MfaFullInfo
	data := utils.Verify(&mfaInfo, c)
	if data == nil {
		data, result := ServiceMfa.MfaCreate(mfaInfo)
		utils.R(data, result, c)
	}
}

// @Tags MFA
// @Description 添加MFA
// @Param account body string true "账号"
// @Param name body string true "名称"
// @Param desc body string false "描述"
// @Param status body int false "所属状态是否有效  1是有效 0是失效"
// @Success 200 {string} string	"ok"
// @Router /api/v1/user/{id} [put]
func MfaUpdate(c *gin.Context) {
	var mfaInfo ModelMfa.MfaFullInfo
	mfaInfo.ID = c.Param("id")
	valid := utils.Verify(&mfaInfo, c)
	if valid == nil {
		data, result := ServiceMfa.MfaUpdate(mfaInfo)
		utils.R(data, result, c)
	}
}

// @Tags MFA
// @Description 删除MFA
// @Param id path string true "id"
// @Router /api/v1/mfa/{id} [delete]
func MfaDel(c *gin.Context) {
	var mfaInfo ModelMfa.MfaInfo
	id := c.Param("id")
	mfaInfo.ID = id
	result := ServiceMfa.MfaDel(mfaInfo)
	utils.R(id, result, c)
}

// Tags MFA
// @Descriptions 获取MFA信息(登录成功之后header携带 token参数获取)
// @Router /api/v1/mfa/getInfo [get]
func GetMfaInfo(c *gin.Context) {
	j := middleware.NewJWT()
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": configure.RequestError,
			"msg":    "请求未携带token，无权限访问",
			"data":   nil,
		})
		c.Abort()
		return
	}
	result, error := j.ParserToken(token)
	if error == nil {
		var mfaInfo ModelMfa.MfaInfo
		mfaInfo.ID = result.ID
		data, r := ServiceMfa.MfaGetInfo(mfaInfo)
		utils.R(data, r, c)
	}
}
