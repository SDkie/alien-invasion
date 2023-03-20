package main

import (
	"log"
	"strconv"

	"github.com/SDkie/alien-invasion/internal/random"
	"github.com/SDkie/alien-invasion/internal/world"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "alien-invasion-cli",
		Short: "alien-invasion-cli allows to run the simulation of alien invasion",
		Args:  cobra.ExactArgs(2),
		Run:   rootCmdRun,
	}

	rootCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		cmd.Println("Usage: alien-invasion-cli <cites-files> <no-of-aliens>")
		return nil
	})

	rootCmd.Execute()
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	fileName := args[0]
	alienCount, err := strconv.Atoi(args[1])
	if err != nil {
		log.Printf("error parsing <no-of-aliens>:%s", err)
		return
	}

	world, err := world.New(fileName, alienCount)
	if err != nil {
		return
	}

	world.Random = random.Random{}
	world.RunAlienInvasion()
}
