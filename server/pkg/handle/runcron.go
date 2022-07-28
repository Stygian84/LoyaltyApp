package handle

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunCron(AccrualTime string, HandbackTime string) {
	s1 := gocron.NewScheduler(time.Local)
	s1.Every(1).Day().At(AccrualTime).Do(SendAccrual)
	s1.StartAsync()

	s2 := gocron.NewScheduler(time.Local)
	s1.Every(1).Day().At(HandbackTime).Do(ReadHandbackFile)
	s2.StartBlocking()
}
