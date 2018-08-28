package cronjobs

import (
	"github.com/robfig/cron"
	"beego_framework/bean"
)

func StartCronjobs(){
	cronjob := cron.New()
	cronjob.AddFunc("@every 2m", func() {
		defer func(){
			if x := recover(); x != nil {
				// just ignore
			}
		}()
		bean.ExchangerServiceBean.UpdateExchangerRate()
	})
	cronjob.Start()
}
