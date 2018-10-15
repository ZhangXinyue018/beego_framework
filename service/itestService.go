package service

import "beego_framework/domain"

type ITestService interface {
	Test() (*[]domain.TestMysql)
}