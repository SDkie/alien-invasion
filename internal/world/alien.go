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
