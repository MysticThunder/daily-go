package task

import (
	"github.com/robfig/cron"
)

func Job(corn *cron.Cron) {
	err := corn.AddFunc("*/15 * * * * *", Push)
	if err != nil {
		return
	}

}
