package world

// Alien struct represents one Alien
type Alien struct {
	Number     int
	MovesCount int
}

// NewAlien creates an Alien using Alien Number
func NewAlien(no int) *Alien {
	return &Alien{
		Number: no,
	}
}

// BuildAliens builds all the aliens required to simulate invasion
func BuildAliens(alienCount int) map[int]*Alien {
	aliens := make(map[int]*Alien)

	for i := 0; i < alienCount; i++ {
		aliens[i] = NewAlien(i)
	}

	return aliens
}
