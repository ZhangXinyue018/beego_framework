package service

import (
	"beego_framework/rpc/thirdparty"
	"fmt"
)

type ExchangerService struct {
	ExchangerRpc *thirdparty.ExchangerRpc
}

func (service *ExchangerService) UpdateExchangerRate() () {
	fmt.Println("updated")
}
