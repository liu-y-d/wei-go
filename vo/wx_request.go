// Package vo @Author yd 2024/3/26 15:57
package vo

type WxLoginRequest struct {
	Code          string `json:"code" form:"code"`
	RawData       string `json:"rawData" form:"rawData"`
	Signature     string `json:"signature" form:"signature"`
	EncryptedData string `json:"encryptedData" form:"encryptedData"`
	Iv            string `json:"iv" form:"iv"`
}
type WxRawData struct {
	Nickname  string `json:"nickname" form:"nickname"`
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl"`
}
