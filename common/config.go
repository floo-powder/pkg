package common

type SyncType string

const (
	// FullSyncType every time sync full data
	FullSyncType SyncType = "Full"
	// IncrementSyncType every time sync increment data
	IncrementSyncType SyncType = "Increment"
)

type ExecuteCycleType string

const (
	// CronExec cron exec job
	CronExec ExecuteCycleType = "Cron"
	// DisposableExec disposable exec job
	DisposableExec ExecuteCycleType = "Disposable"
	// ManualExec manual exec job(never auto exec)
	ManualExec ExecuteCycleType = "Manual"
)
