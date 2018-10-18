package bean

import (
	"github.com/go-xorm/xorm"
	"github.com/robfig/cron"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"gopkg.in/mgo.v2"
)

func InitDatasource() {
	MysqlReadEngineBean = getMysqlEngine(beego.AppConfig.String("read_mysql_url"))
	MysqlWriteEngineBean = getMysqlEngine(beego.AppConfig.String("write_mysql_url"))

	//MongoSessionBean = getMongoSession(beego.AppConfig.String("mongo_url"))
}

func getAdminRedisClient() (*redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("admin_redis_url"),
		Password: beego.AppConfig.String("admin_redis_password"),
		DB:       beego.AppConfig.DefaultInt("admin_redis_db", 0),
	})
	return client
}

func getMongoSession(dbUrl string) (*mgo.Session) {
	MongoSession, err := mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	} else {
		MongoSession.SetMode(mgo.Monotonic, true)
		MongoSession.SetPoolLimit(300)
		return MongoSession
	}
}

func getMysqlEngine(dbUrl string) (*xorm.Engine) {
	engine, err := xorm.NewEngine("mysql", dbUrl)
	if err == nil {
		engine.SetMaxOpenConns(10)
		engine.SetMapper(core.GonicMapper{})
		go pingMysql(engine)
		return engine
	} else {
		panic(err)
	}
}

func pingMysql(engine *xorm.Engine) {
	cronjob := cron.New()
	spec := "*/10 * * * * ?"
	cronjob.AddFunc(spec, func() {
		engine.Ping()
	})
	cronjob.Start()
}
