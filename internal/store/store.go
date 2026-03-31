package store

// Store defines the database operations for the pipeline and API.
type Store interface {
	// IOC operations
	UpsertIOC(record *IOCRecord) error
	GetIOC(indicator string) (*IOCRecord, error)
	GetHighConfidenceIOCs(windowStart, windowEnd interface{}) ([]IOCRecord, error)
	QueryIOCs(iocType, confidence, tag string, limit, offset int) ([]IOCRecord, error)
	DeactivateStaleIOCs() error

	// Summary operations
	InsertCycleSummary(summary *CycleSummary) error
	GetLatestSummary(summaryType string) (*CycleSummary, error)
	GetSummaryByID(id int64) (*CycleSummary, error)
	ListSummaries(limit, offset int) ([]CycleSummary, error)

	// Stub operations
	InsertDetectionStub(stub *DetectionStub) error
	GetStubsByIOC(iocRecordID int64) ([]DetectionStub, error)

	Close() error
}
