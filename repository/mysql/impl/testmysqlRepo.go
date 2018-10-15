package impl

import (
	"beego_framework/domain"
	"github.com/go-xorm/xorm"
)

type TestMysqlRepository struct {
	TestMysqlReadEngine  *xorm.Engine
	TestMysqlWriteEngine *xorm.Engine
}

func (repository *TestMysqlRepository) ListAll() (*[]domain.TestMysql) {
	var testList []domain.TestMysql
	err := repository.TestMysqlReadEngine.Find(&testList)
	if err != nil {
		panic(err)
	}
	return &testList
}
