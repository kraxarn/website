package config

import "github.com/gin-gonic/gin"

func GetDomain() string {
	if gin.Mode() == gin.DebugMode {
		return "localhost"
	}
	return "kraxarn.com"
}

func IsSecure() bool {
	return gin.Mode() != gin.DebugMode
}
