package init

import (
	"hestia/app/controllers"
	depinject "hestia/app/depInject"
	"hestia/internal/database"
	"hestia/internal/logger"
)

func init() {
	initEnv()
	logger.Init()
	defer logger.Close()
	database.InitConnect()
	initMigration()

	// service manager
	container := depinject.NewContainer(database.DB)

	// injection container dans les controllers
	controllers.InitNewsController(container)
}