package impl

import (
	"beego_framework/domain"
	"beego_framework/repository/mysql"
)

type TestService struct {
	Repository mysql.ITestMysqlRepo
}

func (service *TestService) Test() (*[]domain.TestMysql) {
	return service.Repository.ListAll()
}
