package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	*gocron.Scheduler
}

func NewScheduler() (Scheduler, error) {
	location, err := time.LoadLocation(time.FixedZone("UTC", 1).String())
	if err != nil {
		return Scheduler{}, fmt.Errorf("cannot load location: %w", err)
	}

	sch := gocron.NewScheduler(location)

	return Scheduler{sch}, nil
}
