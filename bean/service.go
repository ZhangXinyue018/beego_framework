package bean

import (
	"beego_framework/service"
)

var (
	ExchangerServiceBean *service.ExchangerService
)

func init() {
	ExchangerServiceBean = &service.ExchangerService{
		ExchangerRpc:        ExchangerRpcBean,
	}
}
