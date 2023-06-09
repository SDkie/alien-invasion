package world

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ErrEmptyCitiesFile = errors.New("empty cities file")
)

// City struct represents a City
type City struct {
	Name  string
	Roads map[string]string // direction -> nextCityName
}

// NewCity create a City using city name
func NewCity(name string) *City {
	return &City{
		Name:  name,
		Roads: make(map[string]string),
	}
}

type Cities map[string]*City

// NewCities reads file and creates a map containing all the cities information
func NewCities(fileName string) (Cities, error) {
	cities := make(Cities)

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

		city, err := buildCity(text)
		if err != nil {
			return nil, err
		}

		cities.addCity(city.Name)

		for direction, toCity := range city.Roads {
			cities.addCity(toCity)
			err = cities.addRoad(city.Name, toCity, direction)
			if err != nil {
				return nil, err
			}
		}
	}

	if len(cities) == 0 {
		err := ErrEmptyCitiesFile
		log.Println(err)
		return nil, err
	}

	return cities, nil
}

// addCity adds city if its missing
func (c Cities) addCity(cityName string) {
	_, ok := c[cityName]
	if !ok {
		c[cityName] = NewCity(cityName)
	}
}

// addRoad adds road in both toCity and fromCity whatever is missing
func (c Cities) addRoad(fromCityName, toCityName, direction string) error {
	if fromCityName == toCityName {
		err := fmt.Errorf("invalid city road for %s", toCityName)
		log.Println(err)
		return err
	}

	oppDirection, err := getOppositeDirection(direction)
	if err != nil {
		return err
	}

	fromCity := c[fromCityName]
	toCity := c[toCityName]

	linkCity, ok := fromCity.Roads[direction]
	if !ok {
		fromCity.Roads[direction] = toCityName
	} else if linkCity != toCityName {
		err := fmt.Errorf("invalid city road for %s", fromCityName)
		log.Println(err)
		return err
	}

	linkCity, ok = toCity.Roads[oppDirection]
	if !ok {
		toCity.Roads[oppDirection] = fromCityName
	} else if linkCity != fromCityName {
		err := fmt.Errorf("invalid city road for %s", toCityName)
		log.Println(err)
		return err
	}

	return nil
}

// buildCity build city from the input text string
// text format:'B west=A east=C north=D'
func buildCity(text string) (*City, error) {
	tokens := strings.Split(text, " ")

	cityName := strings.TrimSpace(tokens[0])
	city := NewCity(cityName)

	for i := 1; i < len(tokens); i++ {
		subtokens := strings.Split(tokens[i], "=")
		if len(subtokens) != 2 {
			err := fmt.Errorf("invalid text line in file:%s", text)
			log.Print(err)
			return nil, err
		}

		direction := strings.TrimSpace(subtokens[0])
		nextCity := strings.TrimSpace(subtokens[1])
		city.Roads[direction] = nextCity
	}

	return city, nil
}

// getOppositeDirection takes a direction as input and return the opposite direction
func getOppositeDirection(direction string) (string, error) {
	switch direction {
	case "north":
		return "south", nil
	case "south":
		return "north", nil
	case "east":
		return "west", nil
	case "west":
		return "east", nil
	default:
		err := fmt.Errorf("invalid direction: %s", direction)
		log.Println(err)
		return "", err
	}
}
