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
			aliensInCityFile := fmt.Sprintf("testdata/test%d_aliens_in_city.txt", c.testNo)

			random, aliensCount, err := random.NewMockRandom(alienMoves)
			if err != nil {
				t.Fatal(err)
			}

			w, err := world.New(citiesInput, aliensCount, 10000)
			if err != nil {
				t.Fatal(err)
			}

			w.Random = random
			w.RunAlienInvasion()

			expectedCities, err := world.ReadCitiesFile(citiesOutput)
			if err != nil {
				t.Fatal(err)
			}
			aliensInCity, err := world.ReadAliensInCityFile(aliensInCityFile)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(w.Cities, expectedCities) {
				t.Errorf("w.Cities does not match with %s", citiesOutput)
			}
			if !reflect.DeepEqual(w.AliensInCity, aliensInCity) {
				t.Errorf("w.AliensInCity does not match with %s", aliensInCityFile)
			}
		}
		t.Run(c.name, tf)
	}
}

func TestAliensMoreThanCities(t *testing.T) {
	citiesInput := "testdata/test_aliens_more_than_cities_input.txt"
	alienMoves := "testdata/test_aliens_more_than_cities_aliens.txt"

	_, aliensCount, err := random.NewMockRandom(alienMoves)
	if err != nil {
		t.Fatal(err)
	}
	_, err = world.New(citiesInput, aliensCount, 10000)
	expectedErr := world.ErrAliensMoreThanCities

	if err != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestWithNoAliens(t *testing.T) {
	citiesInput := "testdata/test_with_no_aliens_cities.txt"
	alienMoves := "testdata/test_with_no_aliens_aliens.txt"

	_, aliensCount, err := random.NewMockRandom(alienMoves)
	if err != nil {
		t.Fatal(err)
	}
	_, err = world.New(citiesInput, aliensCount, 10000)
	expectedErr := world.ErrInvalidAliensCount

	if err != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestCityNotFound(t *testing.T) {
	citiesInput := "testdata/test_city_not_found_cities.txt"
	alienMoves := "testdata/test_city_not_found_alien.txt"

	_, aliensCount, err := random.NewMockRandom(alienMoves)
	if err != nil {
		t.Fatal(err)
	}
	_, err = world.New(citiesInput, aliensCount, 10000)
	expectedErr := "B city not found in cities"

	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestInvalidRoadsData(t *testing.T) {
	citiesInput := "testdata/test_invalid_roads_data_cities.txt"
	alienMoves := "testdata/test_invalid_roads_data_aliens.txt"

	_, aliensCount, err := random.NewMockRandom(alienMoves)
	if err != nil {
		t.Fatal(err)
	}
	_, err = world.New(citiesInput, aliensCount, 10000)
	expectedErr := "invalid roads data for city:E"

	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestInvalidLineInCities(t *testing.T) {
	citiesInput := "testdata/test_invalid_line_in_cities_cities.txt"
	alienMoves := "testdata/test_invalid_line_in_cities_aliens.txt"

	_, aliensCount, err := random.NewMockRandom(alienMoves)
	if err != nil {
		t.Fatal(err)
	}
	_, err = world.New(citiesInput, aliensCount, 10000)
	expectedErr := "invalid text line in file:D south=A north"

	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestAlienMaxMovesCount(t *testing.T) {
	citiesInput := "testdata/test_alien_max_moves_count_cities.txt"
	alienCount := 1
	alienMaxMoves := 500

	w, err := world.New(citiesInput, alienCount, alienMaxMoves)
	if err != nil {
		t.Fatal(err)
	}

	w.Random = random.Random{}
	w.RunAlienInvasion()

	if w.MovesCount != alienMaxMoves {
		t.Errorf("MovesCount shoud be:%d got:%d", alienMaxMoves, w.MovesCount)
	}
	if len(w.Aliens) != alienCount {
		t.Errorf("Aliens count should be:%d got:%d", alienCount, len(w.Aliens))
	}
	if len(w.AliensInCity) != alienCount {
		t.Errorf("AliensInCity count should be:%d got:%d", alienCount, len(w.AliensInCity))
	}
}
