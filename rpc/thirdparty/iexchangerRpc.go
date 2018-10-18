package thirdparty

import "beego_framework/domain"

type IExchangerRpc interface {
	ListPrice(symbolNameList *[]string) (*[]domain.SimpleExchangerRate)
}