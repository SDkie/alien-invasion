package world

import (
	"fmt"
	"log"

	"github.com/SDkie/alien-invasion/internal/random"
)

// World struct keep track of the entire world which consists of Cities and Aliens
type World struct {
	Cities map[string]*City // CityName -> *City

	Aliens            map[int]*Alien   // AlienNo -> *Alien
	AliensInCity      map[string][]int // city --> List of AliensNo
	ActiveAliensCount int
	AliensMaxSteps    int

	MovesCount int

	Random random.Randomer
}

// New creates the World instance with cities and aliens
func New(fileName string, aliensCount int, aliensMaxSteps int) (*World, error) {
	var world World
	var err error

	if aliensCount <= 0 {
		err := fmt.Errorf("aliensCount should be greater than 0")
		log.Println(err)
		return nil, err
	}

	world.Cities, err = ReadCitiesFile(fileName)
	if err != nil {
		return nil, err
	}

	if len(world.Cities) < aliensCount {
		err := fmt.Errorf("aliensCount should be less than cities count")
		log.Println(err)
		return nil, err
	}

	world.ActiveAliensCount = aliensCount
	world.AliensMaxSteps = aliensMaxSteps
	world.Aliens = BuildAliens(aliensCount)
	world.AliensInCity = make(map[string][]int)

	return &world, nil
}

// AssignCitiesToAliens assign a city to each Alien
func (w *World) AssignCitiesToAliens() {
	var cities []string
	for cityName := range w.Cities {
		cities = append(cities, cityName)
	}

	var cityName string
	for alienNo := 0; alienNo < w.ActiveAliensCount; alienNo++ {
		// Make sure each city has only one Alien
		for {
			cityName = w.Random.ChooseCity(alienNo, cities)
			_, ok := w.AliensInCity[cityName]
			if !ok {
				break
			}
		}
		w.AliensInCity[cityName] = []int{alienNo}
	}
}

// RunAliensFight check the aliens count in the city and simulate the AliensFight
func (w *World) RunAliensFight() {
	for cityName, aliens := range w.AliensInCity {
		if len(aliens) == 0 {
			delete(w.AliensInCity, cityName)
			continue
		}

		// Single alien so no fight
		if len(aliens) == 1 {
			continue
		}

		log.Printf("\tCity:%s destroyed because of aliens: %v fight\n", cityName, aliens)

		// Remove aliens
		for _, alienNo := range aliens {
			delete(w.Aliens, alienNo)
			w.ActiveAliensCount--
		}

		// Remove roads
		city := w.Cities[cityName]
		for direction, nextCity := range city.Roads {
			delete(w.Cities[nextCity].Roads, getOppositeDirection(direction))
		}

		// Remove city
		delete(w.Cities, cityName)
		delete(w.AliensInCity, cityName)
	}
}

// AliensMove simulate the random moment of the aliens
func (w *World) AliensMove() {
	log.Printf("Aliens Move Count: %d", w.MovesCount)

	for oldCity, aliens := range w.AliensInCity {
		for index, alienNo := range aliens {
			city := w.Cities[oldCity]
			alien := w.Aliens[alienNo]

			// Alien has already moved
			if alien.MovesCount >= w.MovesCount {
				continue
			}

			// Alien is trapped
			if alien.Trapped {
				continue
			}

			// Check if Alien is trapped
			if len(city.Roads) == 0 {
				log.Printf("\tAlien:%d trapped in city:%s", alienNo, city.Name)
				alien.Trapped = true
				w.ActiveAliensCount--
				continue
			}

			var directions []string
			for d := range city.Roads {
				directions = append(directions, d)
			}
			direction := w.Random.ChooseDirection(alienNo, directions)
			newCity := city.Roads[direction]

			w.AliensInCity[oldCity] = aliens[index+1:]
			w.AliensInCity[newCity] = append(w.AliensInCity[newCity], alien.Number)
			alien.MovesCount++
			log.Printf("\tAlien:%d moved %s from %s to %s", alien.Number, direction, oldCity, newCity)
		}
	}
}

// PrintingAllCities prints Cities map
func (w *World) PrintingAllCities() {
	log.Printf("Cities map at the end of simulation:")

	for _, c := range w.Cities {
		msg := c.Name
		for direction, city := range c.Roads {
			msg = fmt.Sprintf("%s %s=%s", msg, direction, city)
		}
		log.Println(msg)
	}
}

// RunAlienInvasion simulates the Alien Invasion
func (w *World) RunAlienInvasion() {
	w.AssignCitiesToAliens()

	for w.MovesCount = 1; w.MovesCount <= w.AliensMaxSteps; w.MovesCount++ {
		if w.ActiveAliensCount == 0 {
			break
		}

		w.AliensMove()
		w.RunAliensFight()
		log.Println()
	}

	w.PrintingAllCities()
}
