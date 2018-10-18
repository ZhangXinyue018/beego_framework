package bean

import (
	"beego_framework/domain/response"
	mysql "beego_framework/repository/mysql/impl"
	thirdparty "beego_framework/rpc/thirdparty/impl"
	"beego_framework/service"
	serviceimpl "beego_framework/service/impl"
	"github.com/go-xorm/xorm"
)

var (
	MysqlReadEngineBean  *xorm.Engine
	MysqlWriteEngineBean *xorm.Engine

	//MongoSessionBean *mgo.Session
)

var ErrorMap map[string]response.BaseResp

var (
	MysqlTempRepoBean *mysql.TempMysqlRepository
)

var (
//MongoTempRepoBean *mongo.TestMongoRepository
)

var (
	ExchangerRpcBean *thirdparty.ExchangerRpc
)

var (
	ExchangerServiceBean *serviceimpl.ExchangerService
	WebSocketServiceBean *service.WebSocketService
	TestServiceBean      *serviceimpl.TestService
)

func init(){
	InitErrorMap()
	InitDatasource()
	InitRepository()
	InitRpc()
	InitService()
}