package init

import (
	"hestia/internal/database"
	"hestia/internal/logger"
)

func init() {
	initEnv()
	logger.Init()
	defer logger.Close()
	database.InitConnect()
	initMigration()
}