package enrichment

// VTClient handles VirusTotal enrichment for HIGH confidence IOCs.
// Enrichment is additive only — VT results do not affect confidence tier.
type VTClient struct {
	APIKey string
}

func NewVTClient(apiKey string) *VTClient {
	return &VTClient{APIKey: apiKey}
}

// Enrich queries VT for the given indicator.
// Returns true if enrichment succeeded, false if rate limited or errored.
// Failures are logged and do not halt the pipeline.
func (c *VTClient) Enrich(indicator string) (bool, error) {
	// TODO: implement VT lookup
	return false, nil
}
