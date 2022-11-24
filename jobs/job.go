package job

import (
	"time"
	h "workspace/handlers"

	"github.com/jasonlvhit/gocron"
)

func RunJob() {
	time.Sleep(50 * time.Second) // Elasticsearch running time
	gocron.Every(4).Week().From(gocron.NextTick()).Do(h.RunSelenium)
	<-gocron.Start()
	s := gocron.NewScheduler()
	<-s.Start()
}
