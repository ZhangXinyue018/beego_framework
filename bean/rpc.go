package bean

import (
	thirdparty "beego_framework/rpc/thirdparty/impl"
	"github.com/astaxie/beego"
)

func InitRpc() {
	ExchangerRpcBean = &thirdparty.ExchangerRpc{
		BitstampUrl:   beego.AppConfig.String("bitstamp_url"),
		CoinMarketUrl: beego.AppConfig.String("coinmarket_url"),
		BitfinexUrl:   beego.AppConfig.String("bitfinex_url"),
	}
}
