package mysql

import (
	"beego_framework/domain"
)

type ITestMysqlRepo interface {
	ListAll() (*[]domain.TestMysql)
}