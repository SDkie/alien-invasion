package world

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// City struct represents a City
type City struct {
	Name  string
	Roads map[string]string // direction -> city
}

// NewCity create a City using city name
func NewCity(name string) *City {
	return &City{
		Name:  name,
		Roads: make(map[string]string),
	}
}

// ReadCitiesFile reads file and creates a map containing all the cities information
func ReadCitiesFile(fileName string) (map[string]*City, error) {
	cities := make(map[string]*City)

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

		cities[city.Name] = city
	}

	return cities, nil
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
