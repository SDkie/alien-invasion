package random

// Randomer is an interface that provide randomness to alien invasion
type Randomer interface {
	ChooseCity(int, []string) string
	ChooseDirection(int, int, []string) string
}
