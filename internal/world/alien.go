package world

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Alien struct represents one Alien
type Alien struct {
	Number     int
	MovesCount int
	Trapped    bool
}

// NewAlien creates an Alien using Alien Number
func NewAlien(no int) *Alien {
	return &Alien{
		Number: no,
	}
}

// BuildAliens builds all the aliens required to simulate invasion
func BuildAliens(alienCount int) map[int]*Alien {
	aliens := make(map[int]*Alien)

	for i := 0; i < alienCount; i++ {
		aliens[i] = NewAlien(i)
	}

	return aliens
}

// ReadAliensInCityFile reads the input file to create the map of cityName to alienNos
func ReadAliensInCityFile(fileName string) (map[string][]int, error) {
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
