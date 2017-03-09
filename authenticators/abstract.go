package authenticators

import "strings"

type AbstractAuthenticator struct {
	DecodedLogin string
	Verified     bool
}

func (aa AbstractAuthenticator) GetAdvertiseAuthentication(authType []string) string {
	if len(authType) == 0 {
		return ""
	}

	return "250-AUTH " + strings.Join(authType, " ") + "\r\n"
}

func (aa AbstractAuthenticator) DecodeLogin() string {
	return aa.DecodedLogin
}

func (aa AbstractAuthenticator) IsVerified() bool {
	return aa.Verified
}
