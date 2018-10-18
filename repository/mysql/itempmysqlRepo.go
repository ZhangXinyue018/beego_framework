package mysql

import (
	"beego_framework/domain"
)

type ITempMysqlRepo interface {
	ListAll() (*[]domain.TestMysql)
}