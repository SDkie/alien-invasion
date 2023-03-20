package world

// World struct keep track of the entire world which consistes of Cities and Aliens
type World struct {
	Cities map[string]*City // CityName -> *City

	Aliens            map[int]*Alien   // AlienNo -> *Alien
	AlienInCity       map[string][]int // city --> List of AliensNo
	ActiveAliensCount int

	MovesCount int
}
