package bean

import (
	"github.com/astaxie/beego"
	"beego_framework/rpc/thirdparty"
)

var (
	ExchangerRpcBean *thirdparty.ExchangerRpc
)

func init() {
	ExchangerRpcBean = &thirdparty.ExchangerRpc{
		BitstampUrl:   beego.AppConfig.String("bitstamp_url"),
		CoinMarketUrl: beego.AppConfig.String("coinmarket_url"),
		BitfinexUrl:   beego.AppConfig.String("bitfinex_url"),
	}
}
