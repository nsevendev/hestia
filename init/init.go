package init

import (
	"hestia/internal/database"
	"hestia/internal/logger"
)

func init() {
	initEnv()
	logger.InitFromEnv()
	defer logger.Close()
	database.Connect()
}