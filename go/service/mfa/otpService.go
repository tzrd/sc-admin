package ServiceMfa

import (
	"scadmin/db"
	"scadmin/utils"
	UtilsDB "scadmin/utils/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ModelMfa "scadmin/model/mfa"
)

func OtpQuery(dates []ModelMfa.OtpInfo, c *gin.Context) ([]ModelMfa.OtpInfo, *gorm.DB) {
	var list []ModelMfa.MfaFullInfo
	var name string = "mfa_info"

	result := db.DB.Scopes(UtilsDB.PaginateByTable(name, c, list)).Find(&list)

	for i := range list {
		mfa := list[i]

		var otpInfo ModelMfa.OtpInfo
		otpInfo.ID = utils.RandSeqStr(32)
		otpInfo.Rid = utils.RandSeqStr(16)
		otpInfo.Issuer = mfa.Issuer
		otpInfo.Account = mfa.Account
		otpInfo.Period = mfa.Period
		otpInfo.Token = utils.GetOtpToken(mfa.Digits, mfa.Secret, mfa.Period)

		dates = append(dates, otpInfo)
	}

	return dates, result
}

func OtpSearch(dates []ModelMfa.OtpInfo, c *gin.Context) ([]ModelMfa.OtpInfo, *gorm.DB) {
	var list []ModelMfa.MfaFullInfo
	var name string = "mfa_info"
	var result *gorm.DB

	q := c.Request.URL.Query().Get("q")
	if len(q) > 0 {
		result = db.DB.Scopes(UtilsDB.PaginateByTable(name, c, list)).Where("account LIKE ?", "%"+q+"%").Find(&list)
	} else {
		result = db.DB.Scopes(UtilsDB.PaginateByTable(name, c, list)).Find(&list)
	}

	for i := range list {
		mfa := list[i]

		var otpInfo ModelMfa.OtpInfo
		otpInfo.ID = utils.RandSeqStr(32)
		otpInfo.Rid = utils.RandSeqStr(16)
		otpInfo.Issuer = mfa.Issuer
		otpInfo.Account = mfa.Account
		otpInfo.Period = mfa.Period
		otpInfo.Token = utils.GetOtpToken(mfa.Digits, mfa.Secret, mfa.Period)

		dates = append(dates, otpInfo)
	}

	return dates, result
}

func OtpSinge(id string) (ModelMfa.OtpInfo, *gorm.DB) {
	var data ModelMfa.OtpInfo
	var mfa ModelMfa.MfaFullInfo
	var name string = "mfa_info"

	result := db.DB.Table(name).First(&mfa, id)

	if result.RowsAffected > 0 {
		data.Rid = utils.RandSeqStr(16)
		data.Issuer = mfa.Issuer
		data.Account = mfa.Account
		data.Period = mfa.Period
		data.Token = utils.GetOtpToken(mfa.Digits, mfa.Secret, mfa.Period)
	}

	return data, result
}
