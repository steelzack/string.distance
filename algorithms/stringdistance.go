package distances

type StringDistance interface {
	CalculateDistance(fromString string, toString string) int
}
