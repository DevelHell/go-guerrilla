package authenticators

import "github.com/DevelHell/go-guerrilla/backends"

type Authenticator interface {
	VerifyLOGIN(login, password string) bool
	//VerifyPLAIN(login, password string) bool
	//VerifyGSSAPI(login, password string) bool
	//VerifyDIGESTMD5(login, password string) bool
	//VerifyMD5(login, password string) bool
	VerifyCRAMMD5(challenge, authString string) bool
	GenerateCRAMMD5Challenge() (string, error)
	ExtractLoginFromAuthString(authString string) string
	DecodeLogin(login string) (string, error)

	GetAdvertiseAuthentication(authType []string) string

	GetMailSize(login string, defaultSize int64) int64
}

type AuthenticatorCreator func(config backends.BackendConfig) Authenticator

type AuthStore struct {
	CRAMMD5challenge string
	IsAuthenticated  bool
}
