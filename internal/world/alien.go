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

type Aliens map[int]*Alien

// NewAliens builds all the aliens required to run invasion
func NewAliens(alienCount int) Aliens {
	aliens := make(Aliens)

	for i := 0; i < alienCount; i++ {
		aliens[i] = NewAlien(i)
	}

	return aliens
}
