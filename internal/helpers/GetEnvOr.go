package helpers

import "os"

func GetEnvOr(envVar, otherwise string) string {
	value, ok := os.LookupEnv(envVar)
	if !ok {
		return otherwise
	}

	return value
}
