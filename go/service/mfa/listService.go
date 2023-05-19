package ServiceMfa

import (
	"scadmin/db"
	UtilsDB "scadmin/utils/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ModelMfa "scadmin/model/mfa"
)

func MfaQuery(datas []ModelMfa.MfaInfo, c *gin.Context) ([]ModelMfa.MfaInfo, *gorm.DB) {
	result := db.DB.Scopes(UtilsDB.Paginate(c, datas)).Find(&datas)
	return datas, result
}

func MfaSearch(datas []ModelMfa.MfaInfo, c *gin.Context) ([]ModelMfa.MfaInfo, *gorm.DB) {
	q := c.Request.URL.Query().Get("q")
	if len(q) > 0 {
		result := db.DB.Scopes(UtilsDB.Paginate(c, datas)).Where("uid LIKE ? OR account LIKE ?", "%"+q+"%", "%"+q+"%").Find(&datas)
		return datas, result
	}

	result := db.DB.Scopes(UtilsDB.Paginate(c, datas)).Find(&datas)
	return datas, result
}

func MfaSinge(id string) (ModelMfa.MfaInfo, *gorm.DB) {
	var data ModelMfa.MfaInfo
	result := db.DB.First(&data, id)
	return data, result
}

func MfaCreate(data ModelMfa.MfaFullInfo) (ModelMfa.MfaFullInfo, *gorm.DB) {
	var name string = "mfa_info"
	data.ID = UtilsDB.CreateUUId()
	result := db.DB.Table(name).Create(&data)
	return data, result
}

func MfaUpdate(data ModelMfa.MfaFullInfo) (ModelMfa.MfaFullInfo, *gorm.DB) {
	var name string = "mfa_info"
	result := db.DB.Table(name).Updates(&data)
	return data, result
}

func MfaDel(data ModelMfa.MfaInfo) *gorm.DB {
	result := db.DB.Delete(&data)
	return result
}

func MfaGetInfo(data ModelMfa.MfaInfo) (ModelMfa.MfaInfo, *gorm.DB) {
	result := db.DB.First(&data)
	return data, result
}
