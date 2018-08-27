package mysql

import "github.com/go-xorm/xorm"

type TestMysqlRepository struct {
	TestMysqlReadEngine  *xorm.Engine
	TestMysqlWriteEngine *xorm.Engine
}

type TestMysql struct {
	Id   int    `xorm:"pk" json:"id"`
	Test string `xorm:"'test'" json:"test"`
}

func (repository *TestMysqlRepository) ListAll() (*[]TestMysql) {
	var testList []TestMysql
	err := repository.TestMysqlReadEngine.Find(&testList)
	if err != nil {
		panic(err)
	}
	return &testList
}
