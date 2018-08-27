package cronjobs

import (
	"github.com/robfig/cron"
	"beego_framework/bean"
)

func CronUpdateExchangerRate(){
	cronjob := cron.New()
	spec := "@every 2m"
	cronjob.AddFunc(spec, func() {
		defer func(){
			if x := recover(); x != nil {
				// just ignore
			}
		}()
		bean.ExchangerServiceBean.UpdateExchangerRate()
	})
	cronjob.Start()
}
