package authenticators

type Authenticator interface {
	VerifyLOGIN(login, password string) bool
	VerifyPLAIN(login, password string) bool
	VerifyGSSAPI(login, password string) bool
	VerifyDIGESTMD5(login, password string) bool
	VerifyMD5(login, password string) bool
	VerifyCRAMMD5(login, password string) bool
	IsVerified() bool
	DecodeLogin() string

	GetAdvertiseAuthentication(authType []string) string
}
