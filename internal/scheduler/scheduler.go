package scheduler

import "time"

// Config holds scheduler timing constants.
type Config struct {
	IngestionInterval time.Duration // default: 1 hour
	ShiftSummaryHours []int         // UTC hours: 6, 14, 22
	DailyResetHour    int           // UTC hour: 12
}

// DefaultConfig returns the schedule defined in the scope.
func DefaultConfig() Config {
	return Config{
		IngestionInterval: 1 * time.Hour,
		ShiftSummaryHours: []int{6, 14, 22},
		DailyResetHour:    12,
	}
}

// Scheduler wires the ingestion ticker and wall clock summary triggers.
type Scheduler struct {
	cfg Config
}

func New(cfg Config) *Scheduler {
	return &Scheduler{cfg: cfg}
}

// Start begins the ingestion ticker and registers summary timers.
// Blocks until ctx is cancelled.
func (s *Scheduler) Start() {
	// TODO: implement ticker + wall clock scheduler
}
