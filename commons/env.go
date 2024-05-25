package commons

import "syscall"

func EnvStrings(key, defaultValue string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return defaultValue
}
