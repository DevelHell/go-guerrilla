package authenticators

import (
	"strings"
)

type AbstractAuthenticator struct{}

func (aa AbstractAuthenticator) GetAdvertiseAuthentication(authType []string) string {
	if len(authType) == 0 {
		return ""
	}

	return "250-AUTH " + strings.Join(authType, " ") + "\r\n"
}

func (aa AbstractAuthenticator) GetMailSize(login string, defaultSize int64) int64 {
	return defaultSize
}
