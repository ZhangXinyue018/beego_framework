package impl

import (
	"beego_framework/domain"
	"github.com/go-xorm/xorm"
)

type TempMysqlRepository struct {
	TempMysqlReadEngine  *xorm.Engine
	TempMysqlWriteEngine *xorm.Engine
}

func (repository *TempMysqlRepository) ListAll() (*[]domain.TestMysql) {
	var testList []domain.TestMysql
	err := repository.TempMysqlReadEngine.Find(&testList)
	if err != nil {
		panic(err)
	}
	return &testList
}
