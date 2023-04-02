package types

const (
	QRCODE_STATE_NEW   = "NEW:"
	QRCODE_STATE_SCAN  = "SCAN:"
	QRCODE_STATE_LOGIN = "LOGIN:"
)

type QrCodeResp struct {
	State string `json:"state"`
	User  UserInfo
}
type UserInfo struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}
