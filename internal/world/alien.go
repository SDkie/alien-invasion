package world

// Alien struct represents an Alien
type Alien struct {
	Number     int
	MovesCount int
	Trapped    bool
}

// NewAlien creates an Alien using Alien Number
func NewAlien(no int) *Alien {
	return &Alien{
		Number: no,
	}
}

// BuildAliens builds all the aliens required to run invasion
func BuildAliens(alienCount int) map[int]*Alien {
	aliens := make(map[int]*Alien)

	for i := 0; i < alienCount; i++ {
		aliens[i] = NewAlien(i)
	}

	return aliens
}
