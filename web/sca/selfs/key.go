package selfs

import "strings"

type SetType int

const (
	Env SetType = iota
	Netrc
	Unknown
)

func KeyClassify(key string) SetType {
	key = strings.Split(key, "=")[0]
	if strings.HasPrefix(key, "CLADMIN_") {
		return Env
	} else if strings.Contains(key, ".") {
		return Netrc
	}
	return Unknown
}
