package env

import (
	"os"
	"strings"
)

const (
	XEnvKey     = "X_ENV"
	XEnvOnline  = "online"
	XEnvOffline = "offline"
)

func GetEnv() string {
	val := os.Getenv(XEnvKey)
	if len(val) == 0 {
		return XEnvOffline
	}
	if strings.ToLower(val) == XEnvOnline {
		return XEnvOnline
	}
	return XEnvOffline
}

func IsOffline() bool {
	return GetEnv() == XEnvOffline
}

func IsOnline() bool {
	return GetEnv() == XEnvOnline
}
