package bean

import (
	mysql "beego_framework/repository/mysql/impl"
	)

var (
	MysqlTestRepoBean *mysql.TestMysqlRepository
)

var (
	//MongoTestRepoBean *mongo.TestMongoRepository
)

func init() {
	MysqlTestRepoBean = &mysql.TestMysqlRepository{
		TestMysqlReadEngine:  MysqlReadEngineBean,
		TestMysqlWriteEngine: MysqlWriteEngineBean,
	}

	//MongoTestRepoBean = &mongo.TestMongoRepository{
	//	TestMongoSession:   MongoSessionBean,
	//	TestDBName:         "test",
	//	TestCollectionName: "test",
	//}
}
