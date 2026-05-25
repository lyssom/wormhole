package scheduler

import (
	"fmt"
	"log"
	"sync"
	"time"

	"wormhole/server/storage"
)

// SchedulerOptions configures the MergeScheduler.
type SchedulerOptions struct {
	Interval       time.Duration // Scan interval, default 30s
	MergeThreshold int          // Segment count threshold to trigger merge, default 3
}

// DefaultSchedulerOptions returns default scheduler options.
func DefaultSchedulerOptions() *SchedulerOptions {
	return &SchedulerOptions{
		Interval:       30 * time.Second,
		MergeThreshold: 3,
	}
}

// MergeScheduler runs background merge operations on immutable segments.
type MergeScheduler struct {
	engine   *storage.StorageEngine
	opts     *SchedulerOptions
	ticker   *time.Ticker
	stopCh   chan struct{}
	stopDone chan struct{}
	mu       sync.Mutex
	stopped  bool
}

// NewMergeScheduler creates a new MergeScheduler.
func NewMergeScheduler(engine *storage.StorageEngine, opts *SchedulerOptions) *MergeScheduler {
	if opts == nil {
		opts = DefaultSchedulerOptions()
	}
	return &MergeScheduler{
		engine:   engine,
		opts:     opts,
		stopCh:   make(chan struct{}),
		stopDone: make(chan struct{}),
	}
}

// Start begins the background ticker loop.
func (s *MergeScheduler) Start() {
	s.mu.Lock()
	if s.ticker != nil {
		s.mu.Unlock()
		return
	}
	s.ticker = time.NewTicker(s.opts.Interval)
	s.mu.Unlock()

	go s.run()
}

// Stop gracefully stops the scheduler.
func (s *MergeScheduler) Stop() error {
	s.mu.Lock()
	if s.stopped {
		s.mu.Unlock()
		return nil
	}
	s.stopped = true
	close(s.stopCh)
	s.mu.Unlock()

	<-s.stopDone
	return nil
}

func (s *MergeScheduler) run() {
	defer close(s.stopDone)
	for {
		// Check stopped flag inside lock
		s.mu.Lock()
		if s.stopped {
			s.mu.Unlock()
			return
		}
		s.mu.Unlock()

		select {
		case <-s.ticker.C:
			s.scanAndMerge()
		case <-s.stopCh:
			return
		}
	}
}

// scanAndMerge scans all tables and triggers merge for partitions exceeding threshold.
func (s *MergeScheduler) scanAndMerge() {
	tables := s.engine.Tables()
	for _, table := range tables {
		if err := s.checkAndMergeTable(table); err != nil {
			log.Printf("scanAndMerge: table %s: %v", table.Name(), err)
		}
	}
}

// checkAndMergeTable checks if any partition needs merging and triggers merge.
func (s *MergeScheduler) checkAndMergeTable(table *storage.Table) error {
	// For now, use SchedulerSegments which groups by "default" partition
	// After Task 3 (Partition), this will use proper partition-aware grouping
	partitions := table.SchedulerSegments()
	for partKey, segs := range partitions {
		if len(segs) >= s.opts.MergeThreshold {
			task := NewMergeTask(table.Name(), table.DB(), partKey, segs, s.engine.RootDir())
			if err := task.Execute(); err != nil {
				return fmt.Errorf("merge task failed: %w", err)
			}
			// Replace old segments with new merged one
			if err := table.ReplaceImmutableSegments(task.InputPaths(), task.OutputPath()); err != nil {
				return fmt.Errorf("ReplaceImmutableSegments failed: %w", err)
			}
		}
	}
	return nil
}

// SegmentInfo holds information about a segment for merge planning.
type SegmentInfo struct {
	Path     string
	RowCount uint64
	TsMin    uint64
	TsMax    uint64
}
