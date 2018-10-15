package cronjobs

import (
	"beego_framework/bean"
	"github.com/robfig/cron"
)

func StartCronjobs() {
	cronjob := cron.New()
	cronjob.AddFunc("@every 30s", func() {
		bean.ExchangerServiceBean.UpdateExchangerRate()
	})
	cronjob.Start()
}
