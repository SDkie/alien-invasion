package random

import "math/rand"

// Random implements Randomer using rand function
type Random struct {
}

// ChooseCity choose randomly one city for alienNo
func (r Random) ChooseCity(alienNo int, cities []string) string {
	cityNo := rand.Intn(len(cities))
	return cities[cityNo]
}

// ChooseDirection choose randomly one direction for alienNo
func (r Random) ChooseDirection(alienNo int, moveCount int, directions []string) string {
	n := rand.Intn(len(directions))
	return directions[n]
}
