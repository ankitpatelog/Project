package config

import "os"

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret=="" {
		panic("JWT_SECRET not set")
	}
	return secret
}