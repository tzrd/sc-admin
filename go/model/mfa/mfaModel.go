package ModelMfa

import "scadmin/model"

// MFA
type MfaInfo struct {
	model.CommonCreate
	Uid     string `json:"uid"`
	Issuer  string `json:"issuer"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Status  int    `json:"status"`
}

// MFA Full
type MfaFullInfo struct {
	model.CommonCreate
	Uid     string `json:"uid"`
	Issuer  string `json:"issuer"`
	Account string `json:"account"`
	Digits  int    `json:"digits"`
	Period  int    `json:"period"`
	Qrcode  string `json:"qrCode"`
	Secret  string `json:"secret"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Status  int    `json:"status"`
}

// OTP
type OtpInfo struct {
	ID      string `json:"id" gorm:"primary_key"`
	Rid     string `json:"rid"`
	Issuer  string `json:"issuer"`
	Account string `json:"account"`
	Period  int    `json:"period"`
	Token   string `json:"token"`
}
