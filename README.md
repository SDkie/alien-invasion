# alien-invasion

---
* [Intro](#intro)
* [Build](#build)
* [Run](#run)
* [Testing](#testing)
---
## Intro:
* first parameter is cities file similar to `sample_cities.txt` in the repo
* second parameter is number of aliens
* `<no-of-aliens>` should not be greater than cities count
* Alien Max Steps is 10000 by default, but it can be changed using `--alien-max-steps` flag
* Program execution stops if all the aliens are either dead/trapped or reach `alien-max-steps` count

## Build
    go build ./cmd/alien-invasion-cli

## Run
Basic cli command

    alien-invasion-cli <cites-file> <no-of-aliens>
    
Using optional flag

    alien-invasion-cli <cites-file> <no-of-aliens> --alien-max-steps <steps-count>

Sample command

    alien-invasion-cli sample_cities.txt 3

## Testing
    go test ./...
