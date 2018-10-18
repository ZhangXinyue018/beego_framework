package bean

import (
	mysql "beego_framework/repository/mysql/impl"
)

func InitRepository() {
	MysqlTempRepoBean = &mysql.TempMysqlRepository{
		TempMysqlReadEngine:  MysqlReadEngineBean,
		TempMysqlWriteEngine: MysqlWriteEngineBean,
	}

	//MongoTestRepoBean = &mongo.TestMongoRepository{
	//	TempMongoSession:   MongoSessionBean,
	//	TempDBName:         "test",
	//	TempCollectionName: "test",
	//}
}
