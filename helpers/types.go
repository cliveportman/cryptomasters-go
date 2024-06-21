package helpers

type Result struct {
	Character string
	Score     int
	Text      string
}

type HammingDistanceResult struct {
	KeySize int
	Score   float64
}