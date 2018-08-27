package bean

import (
	"beego_framework/repository/mysql"
	"beego_framework/repository/mongo"
)

var (
	MysqlTestRepoBean *mysql.TestMysqlRepository
)

var (
	MongoTestRepoBean *mongo.TestMongoRepository
)

func init() {
	MysqlTestRepoBean = &mysql.TestMysqlRepository{
		TestMysqlReadEngine:  MysqlReadEngineBean,
		TestMysqlWriteEngine: MysqlWriteEngineBean,
	}

	MongoTestRepoBean = &mongo.TestMongoRepository{
		TestMongoSession:   MongoSessionBean,
		TestDBName:         "test",
		TestCollectionName: "test",
	}
}
