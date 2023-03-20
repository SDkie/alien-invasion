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
// startCity defines the first city of the aliens
// directions defiends all the directions of the aliens
func NewMockRandom(fileName string) (MockRandom, error) {
	var random MockRandom

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("error opening file:%s\n", err)
		return random, err
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

	return random, nil
}

func (r MockRandom) ChooseCity(alienNo int, cities []string) string {
	return r.startCity[alienNo]
}

func (r MockRandom) ChooseDirection(alienNo int, moveCount int, directions []string) string {
	return r.directions[alienNo][moveCount]
}
