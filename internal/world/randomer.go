package world

// Randomer is an interface that provide functions to choose city and direction
type Randomer interface {
	ChooseCity(int, []string) string
	ChooseDirection(int, []string) string
}
