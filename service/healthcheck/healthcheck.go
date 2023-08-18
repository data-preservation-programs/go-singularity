package healthcheck

import (
	"context"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var staleThreshold = time.Minute * 5
var reportInterval = time.Minute

var cleanupInterval = time.Minute * 5

type State struct {
	WorkType  model.WorkType
	WorkingOn string
}

var logger = log.Logger("healthcheck")

// StartHealthCheckCleanup continuously runs the HealthCheckCleanup function
// at intervals specified by the cleanupInterval. The function cleans up resources
// or updates the health status of various components in your application, as
// implemented by HealthCheckCleanup.
//
// It is designed to be run as a background task and will continue to run
// until the passed context is cancelled.
//
// Parameters:
//   - ctx context.Context: The context that controls cancellations and timeouts.
//     When the context is cancelled (e.g. during shutdown), the function returns,
//     stopping its background cleaning task.
//   - db *gorm.DB: The database connection object used by HealthCheckCleanup to interact with
//     the database.
func StartHealthCheckCleanup(ctx context.Context, db *gorm.DB) {
	for {
		HealthCheckCleanup(ctx, db)
		select {
		case <-ctx.Done():
			return
		case <-time.After(cleanupInterval):
			continue
		}
	}
}

// HealthCheckCleanup is a function that cleans up stale workers and work files in the database.
//
// It first removes all workers that haven't sent a heartbeat for a certain threshold (staleThreshold).
// If there's an error removing the workers, it logs the error and continues.
//
// Then, it resets the state of any sources that are marked as being processed by a worker that no longer exists.
// If there's an error updating the sources, it logs the error and continues.
//
// Finally, it resets the state of any pack jobs that are marked as being packed by a worker that no longer exists.
// If there's an error updating the pack jobs, it logs the error.
//
// All database operations are retried on failure using the DoRetry function.
//
// Parameters:
// db: The Gorm DBNoContext connection to use for database queries.
func HealthCheckCleanup(ctx context.Context, db *gorm.DB) {
	db = db.WithContext(ctx)
	logger.Debugw("running healthcheck cleanup")
	// Remove all workers that haven't sent heartbeat for 5 minutes
	err := database.DoRetry(ctx, func() error {
		return db.Where("last_heartbeat < ?", time.Now().UTC().Add(-staleThreshold)).Delete(&model.Worker{}).Error
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Logger("healthcheck").Errorw("failed to remove dead workers", "error", err)
	}

	// In case there are some works that have stale foreign key referenced to dead workers, we need to remove them
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Source{}).Where("(dag_gen_worker_id NOT IN (?) OR dag_gen_worker_id IS NULL) AND dag_gen_state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"dag_gen_worker_id": nil,
				"dag_gen_state":     model.Ready,
			}).Error
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Logger("healthcheck").Errorw("failed to remove stale daggen worker", "error", err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Source{}).Where("(scanning_worker_id NOT IN (?) OR scanning_worker_id IS NULL) AND scanning_state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"scanning_worker_id": nil,
				"scanning_state":     model.Ready,
			}).Error
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Logger("healthcheck").Errorw("failed to remove stale scanning worker", "error", err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.PackJob{}).Where("(packing_worker_id NOT IN (?) OR packing_worker_id IS NULL) AND packing_state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"packing_worker_id": nil,
				"packing_state":     model.Ready,
			}).Error
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Logger("healthcheck").Errorw("failed to remove stale packing worker", "error", err)
	}
}

// Register registers a new worker in the database. It uses the provided context and database connection.
// The workerID is used to uniquely identify the worker. The getState function is used to get the current state of the worker.
// If allowDuplicate is set to true, it allows the registration of duplicate workers.
//
// The function returns two values:
// - alreadyRunning: A boolean indicating if the worker is already running.
// - err: An error that will be nil if no errors occurred.
//
// The function first gets the hostname of the machine where it's running. If it fails to get the hostname, it returns an error.
// Then it gets the current state of the worker using the getState function.
// It then creates a new worker model with the provided workerID, the current time as the last heartbeat, the hostname, and the work type and working on values from the state.
//
// If allowDuplicate is set to false, the function checks if there are any active workers with the same work type and whose last heartbeat is not stale.
// If there are such workers, it sets alreadyRunning to true and returns.
//
// Finally, it tries to create the worker in the database. If it fails, it returns an error.
func Register(ctx context.Context, db *gorm.DB, workerID uuid.UUID, getState func() State, allowDuplicate bool) (alreadyRunning bool, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return false, errors.Wrap(err, "failed to get hostname")
	}
	state := getState()
	worker := model.Worker{
		ID:            workerID.String(),
		LastHeartbeat: time.Now().UTC(),
		Hostname:      hostname,
		WorkType:      state.WorkType,
		WorkingOn:     state.WorkingOn,
	}
	logger.Debugw("registering worker", "worker", worker)
	err = database.DoRetry(ctx, func() error {
		if !allowDuplicate {
			var activeWorkerCount int64
			err := db.WithContext(ctx).Model(&model.Worker{}).Where("work_type = ? AND last_heartbeat > ?", state.WorkType, time.Now().UTC().Add(-staleThreshold)).
				Count(&activeWorkerCount).Error
			if err != nil {
				return errors.Wrap(err, "failed to count active workers")
			}
			if activeWorkerCount > 0 {
				alreadyRunning = true
				return nil
			}
		}

		return db.WithContext(ctx).Create(&worker).Error
	})

	return alreadyRunning, errors.WithStack(err)
}

// ReportHealth reports the health of a worker to the database. It uses the provided context and database connection.
// The workerID is used to uniquely identify the worker. The getState function is used to get the current state of the worker.
//
// The function first gets the hostname of the machine where it's running. If it fails to get the hostname, it logs an error and returns.
// Then it gets the current state of the worker using the getState function.
// It then creates a new worker model with the provided workerID, the current time as the last heartbeat, the hostname, and the work type and working on values from the state.
//
// The function then tries to create the worker in the database or update the existing worker if one with the same ID already exists.
// The update will set the last heartbeat, work type, working on, and hostname fields to the values from the worker model.
// If the database operation fails, it logs an error.
func ReportHealth(ctx context.Context, db *gorm.DB, workerID uuid.UUID, getState func() State) {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Errorw("failed to get hostname", "error", err)
		return
	}
	state := getState()
	worker := model.Worker{
		ID:            workerID.String(),
		LastHeartbeat: time.Now().UTC(),
		Hostname:      hostname,
		WorkType:      state.WorkType,
		WorkingOn:     state.WorkingOn,
	}
	logger.Debugw("sending heartbeat", "worker", worker)
	err = database.DoRetry(ctx, func() error {
		return db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"last_heartbeat", "work_type", "working_on", "hostname"}),
		}).Create(&worker).Error
	})

	if err != nil {
		logger.Errorw("failed to send heartbeat", "error", err)
	}
}

// StartReportHealth continuously runs the ReportHealth function at intervals
// specified by the reportInterval. This function is responsible for reporting
// the health status of a worker to a centralized store (e.g., a database).
//
// The health status of the worker is determined by calling the provided getState
// function, which should return the current state of the worker.
//
// This function is designed to be run as a background task and will continue to
// run until the passed context is cancelled.
//
// Parameters:
//   - ctx context.Context: The context that controls cancellations and timeouts.
//     When the context is cancelled (e.g. during shutdown), the function returns,
//     stopping its background health reporting task.
//   - db *gorm.DB: The database connection object used by ReportHealth to interact with
//     the database.
//   - workerID uuid.UUID: The unique identifier for the worker whose health is being reported.
//   - getState func() State: A function that, when called, returns the current health
//     state of the worker. The returned State object is expected to contain the
//     necessary information for the ReportHealth function.
func StartReportHealth(ctx context.Context, db *gorm.DB, workerID uuid.UUID, getState func() State) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(reportInterval):
			ReportHealth(ctx, db, workerID, getState)
		}
	}
}
