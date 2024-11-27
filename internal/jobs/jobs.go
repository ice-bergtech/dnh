package jobs

import (
	"fmt"
	"log/slog"

	"github.com/go-co-op/gocron/v2"
	"github.com/ice-bergtech/dnh/src/internal/model_ent"
)

func Init() *gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		slog.Error("Failed to create scheduler", "error", err)
		return nil
	}
	return &s
}

func AddJob(s gocron.Scheduler, job model_ent.ScanJob) error {
	if s == nil {
		return fmt.Errorf("scheduler is nil")
	}
	s.NewJob(
		gocron.OneTimeJob(gocron.OneTimeJobStartImmediately()),
		gocron.NewTask(job.FuncPtr, job.FuncParams...),
	)
	return nil
}
