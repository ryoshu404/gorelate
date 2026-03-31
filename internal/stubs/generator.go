package stubs

import "github.com/ryoshu404/gorelate/internal/store"

// Generate produces a YARA detection stub for a HIGH confidence IOC.
// Template selection is based on indicator type (ip, domain, url, hash).
func Generate(record *store.IOCRecord) (string, error) {
	// TODO: implement template-based YARA generation
	return "", nil
}
