package fetchers

// URLHausFetcher fetches indicators from URLHaus.
type URLHausFetcher struct{}

func NewURLHausFetcher() *URLHausFetcher {
	return &URLHausFetcher{}
}

func (f *URLHausFetcher) Name() string { return "urlhaus" }

func (f *URLHausFetcher) Fetch() ([]IOCRecord, error) {
	// TODO: implement URLHaus feed ingestion
	return nil, nil
}
