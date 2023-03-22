package random

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// MockRandom implements Randomer for testing purpose
//
// startCities and directions of the aliens are pre-defined in a file
type MockRandom struct {
	startCity  []string
	directions [][]string
}

// NewMockRandom reads the fileName and initializes the MockRandom struct
//
// startCity defines the starting city of the aliens
// directions defines all the directions which aliens takes
// It also returns the alienCount which is required to run the AlienInvasion
func NewMockRandom(fileName string) (MockRandom, int, error) {
	var random MockRandom

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("error opening file:%s\n", err)
		return random, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		tokens := strings.Split(text, " ")
		random.startCity = append(random.startCity, strings.TrimSpace(tokens[0]))
		random.directions = append(random.directions, tokens[1:])
	}

	alienCount := len(random.startCity)

	return random, alienCount, nil
}

func (r MockRandom) ChooseCity(alienNo int, cities []string) string {
	return r.startCity[alienNo]
}

func (r MockRandom) ChooseDirection(alienNo int, directions []string) string {
	direction := r.directions[alienNo][0]
	r.directions[alienNo] = r.directions[alienNo][1:]

	return direction
}
