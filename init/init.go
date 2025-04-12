package init

import (
	depinject "hestia/app/depInject"
	"hestia/internal/database"
	"hestia/internal/logger"
)

var Container *depinject.Container

/*
initialisation des modules
environnement, logger, database, migration, container de service
*/
func Bootstrap() {
	initEnv()
	logger.Init()
	defer logger.Close()
	database.InitConnect()
	initMigration()
	Container = depinject.NewContainer(database.DB)
}