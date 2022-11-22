package cron

import (
	"time"

	"github.com/Abeldlp/fullinfo/scheduler-service/controller"
	"github.com/go-co-op/gocron"
)

func ScheduleSaveSateliteLocation(amountOfSeconds int) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(amountOfSeconds).Seconds().Do(func() {
		controller.SaveSateliteLocation()
	})

	s.StartBlocking()
}
