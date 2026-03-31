package scorer

// Confidence tiers.
const (
	Low    = "low"
	Medium = "medium"
	High   = "high"
)

// Score returns a confidence tier based on source count.
//
//	1 source  → low
//	2 sources → medium
//	3 sources → high
func Score(sourceCount int) string {
	switch {
	case sourceCount >= 3:
		return High
	case sourceCount == 2:
		return Medium
	default:
		return Low
	}
}
