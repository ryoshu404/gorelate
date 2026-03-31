package fetchers

// AbuseCHFetcher fetches indicators from abuse.ch feeds.
type AbuseCHFetcher struct{}

func NewAbuseCHFetcher() *AbuseCHFetcher {
	return &AbuseCHFetcher{}
}

func (f *AbuseCHFetcher) Name() string { return "abuse.ch" }

func (f *AbuseCHFetcher) Fetch() ([]IOCRecord, error) {
	// TODO: implement abuse.ch feed ingestion
	return nil, nil
}
