package fetchers

// OTXFetcher fetches indicators from AlienVault OTX.
type OTXFetcher struct {
	APIKey string
}

func NewOTXFetcher(apiKey string) *OTXFetcher {
	return &OTXFetcher{APIKey: apiKey}
}

func (f *OTXFetcher) Name() string { return "otx" }

func (f *OTXFetcher) Fetch() ([]IOCRecord, error) {
	// TODO: implement OTX pulse feed ingestion
	return nil, nil
}
