package impl

import (
	"beego_framework/domain"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
	"testing"
)

var (
	MysqlDBType  = "sqlite3"
	MysqlConnStr = "./test.db?cache=shared&mode=rwc"
	MysqlFile    = "./test.db"
)

func StartMysql(t *testing.T) (*xorm.Engine) {
	os.RemoveAll(MysqlFile)
	testEngine, err := xorm.NewEngine(MysqlDBType, MysqlConnStr)
	if err != nil {
		t.Error(err.Error())
	}
	testEngine.ShowSQL(true)
	testEngine.SetLogLevel(core.LOG_DEBUG)
	testEngine.SetMapper(core.SnakeMapper{})
	initData(testEngine, t)
	return testEngine
}

func StopMysql(engine *xorm.Engine) () {
	engine.Close()
	os.RemoveAll(MysqlFile)
}

func initData(engine *xorm.Engine, t *testing.T) {
	err := engine.CreateTables(&domain.TestMysql{})
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
