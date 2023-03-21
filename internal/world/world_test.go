package world_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/SDkie/alien-invasion/internal/random"
	"github.com/SDkie/alien-invasion/internal/world"
)

func TestRunAlienInvasion(t *testing.T) {
	cases := []struct {
		name   string
		testNo int
	}{
		{name: "City A destroyed", testNo: 1},
		{name: "City D destroyed", testNo: 2},
		{name: "City E destroyed", testNo: 3},
		{name: "Cities A and E destroyed", testNo: 4},
		{name: "No Cities destroyed", testNo: 5},
		{name: "One Alien trapped", testNo: 6},
		{name: "Two Alien trapped", testNo: 7},
	}

	for _, c := range cases {
		tf := func(t *testing.T) {
			citiesInput := fmt.Sprintf("testdata/test%d_cities_input.txt", c.testNo)
			citiesOutput := fmt.Sprintf("testdata/test%d_cities_output.txt", c.testNo)
			alienMoves := fmt.Sprintf("testdata/test%d_alien_moves.txt", c.testNo)

			random, aliensCount, err := random.NewMockRandom(alienMoves)
			if err != nil {
				t.Error(err)
			}

			w, err := world.New(citiesInput, aliensCount, 10000)
			if err != nil {
				t.Error(err)
			}

			w.Random = random
			w.RunAlienInvasion()

			expectedCities, err := world.ReadCitiesFile(citiesOutput)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(w.Cities, expectedCities) {
				t.Errorf("w.Cities does not match with %s", citiesOutput)
			}
		}
		t.Run(c.name, tf)
	}
}
