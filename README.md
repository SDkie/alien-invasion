# alien-invasion

---
* [Assumptions](#assumptions)
* [Build](#build)
* [Usage](#usage)
* [Testing](#testing)
---
## Assumptions:
* Aliens count should not be greater than cities count

## Build
    go build ./cmd/alien-invasion-cli

## Usage


    alien-invasion-cli allows to run the simulation of alien invasion
    
    Usage:
      alien-invasion-cli <cites-file> <no-of-aliens> [flags]
    
    Examples:
      alien-invasion-cli sample_cities.txt 2
    
    Flags:
      -s, --alien-max-steps int   set max steps for aliens (default 10000)
      -h, --help                  help for alien-invasion-cli

## Testing
    go test ./...
