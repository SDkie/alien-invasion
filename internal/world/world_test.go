package world_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
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
			alienMaxMoves := 10000

			random, aliensCount, err := random.NewMockRandom(alienMoves)
			if err != nil {
				t.Fatal(err)
			}

			w, err := world.New(citiesInput, aliensCount, alienMaxMoves)
			if err != nil {
				t.Fatal(err)
			}

			w.Random = random
			w.RunAlienInvasion()

			expectedCities, err := world.NewCities(citiesOutput)
			if err != nil {
				t.Fatal(err)
			}
			aliensInCity, err := readAliensInCityFile(aliensInCityFile)
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
	// NOTE: cities File has 6 cities
	citiesInput := "testdata/test_aliens_more_than_cities_input.txt"
	alienCount := 7
	alienMaxMoves := 10000
	expectedErr := world.ErrAliensMoreThanCities

	_, err := world.New(citiesInput, alienCount, alienMaxMoves)
	if err != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestWithNoAliens(t *testing.T) {
	citiesInput := "testdata/test_with_no_aliens_cities.txt"
	alienCount := 0
	alienMaxMoves := 10000
	expectedErr := world.ErrInvalidAliensCount

	_, err := world.New(citiesInput, alienCount, alienMaxMoves)
	if err != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestInvalidLineInCities(t *testing.T) {
	// NOTE: Invalid Line in file: 'D south=A north'
	citiesInput := "testdata/test_invalid_line_in_cities_cities.txt"
	alienCount := 2
	alienMaxMoves := 10000
	expectedErr := "invalid text line in file:D south=A north"

	_, err := world.New(citiesInput, alienCount, alienMaxMoves)
	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestInvalidCityRoad(t *testing.T) {
	// NOTE: test file has invalid road links for city E
	citiesInput := "testdata/test_invalid_city_road_cities.txt"
	alienCount := 2
	alienMaxMoves := 10000
	expectedErr := "invalid city road for E"

	_, err := world.New(citiesInput, alienCount, alienMaxMoves)
	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestInvalidDirectionInCities(t *testing.T) {
	// The direction text 'eaast' is invalid in the cities file
	citiesInput := "testdata/test_invalid_direction_in_cities.txt"
	alienCount := 2
	alienMaxMoves := 10000
	expectedErr := "invalid direction: eaast"

	_, err := world.New(citiesInput, alienCount, alienMaxMoves)
	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected err:'%s' got:'%s'", expectedErr, err)
	}
}

func TestRoadToSameCity(t *testing.T) {
	// NOTE: The city F is linked to itself
	citiesInput := "testdata/test_road_to_same_city.txt"
	alienCount := 2
	alienMaxMoves := 10000
	expectedErr := "invalid city road for F"

	_, err := world.New(citiesInput, alienCount, alienMaxMoves)
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

	if w.CurrAliensMovesCount != alienMaxMoves {
		t.Errorf("MovesCount shoud be:%d got:%d", alienMaxMoves, w.CurrAliensMovesCount)
	}
	if len(w.Aliens) != alienCount {
		t.Errorf("Aliens count should be:%d got:%d", alienCount, len(w.Aliens))
	}
	if len(w.AliensInCity) != alienCount {
		t.Errorf("AliensInCity count should be:%d got:%d", alienCount, len(w.AliensInCity))
	}
}

// readAliensInCityFile reads the input file to create the map of cityName to alienNos
func readAliensInCityFile(fileName string) (map[string][]int, error) {
	aliensInCity := make(map[string][]int)

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("error opening file:%s", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		tokens := strings.Split(text, "=")
		if len(tokens) != 2 {
			err := fmt.Errorf("invalid text line in file:%s", text)
			log.Print(err)
			return nil, err
		}

		cityName := strings.TrimSpace(tokens[0])
		alienNo, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Printf("error parsing alienNo:%s", err)
			return nil, err
		}
		aliensInCity[cityName] = []int{alienNo}
	}

	return aliensInCity, nil
}
