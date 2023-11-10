package handle

import (
	"time"

	"github.com/go-co-op/gocron"
)

// The time is in 24 hrs format. AccrualTime is the time to send accrual data to sftp server, while HandbackTime is the time to receive handback file from sftp server.
func RunCron(AccrualTime string, HandbackTime string) {
	s1 := gocron.NewScheduler(time.Local)
	s1.Every(1).Day().At(AccrualTime).Do(SendAccrual)
	s1.StartAsync()

	s2 := gocron.NewScheduler(time.Local)
	s1.Every(1).Day().At(HandbackTime).Do(ReadHandbackFile)
	s2.StartBlocking()
}
