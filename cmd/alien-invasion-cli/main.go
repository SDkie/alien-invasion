package main

import (
	"log"
	"strconv"

	"github.com/SDkie/alien-invasion/internal/random"
	"github.com/SDkie/alien-invasion/internal/world"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)

	rootCmd := &cobra.Command{
		Use:     "alien-invasion-cli <cites-file> <no-of-aliens>",
		Short:   "alien-invasion-cli allows to run the simulation of alien invasion",
		Example: "  alien-invasion-cli sample_cities.txt 2",
		Args:    cobra.ExactArgs(2),
		Run:     rootCmdRun,
	}
	rootCmd.PersistentFlags().IntP("alien-max-steps", "s", 10000, "set max steps for aliens")

	rootCmd.Execute()
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	alienMaxSteps, err := cmd.Flags().GetInt("alien-max-steps")
	if err != nil {
		log.Printf("error getting alien-max-steps flag:%s", err)
		return
	}

	fileName := args[0]
	alienCount, err := strconv.Atoi(args[1])
	if err != nil {
		log.Printf("error parsing <no-of-aliens>:%s", err)
		return
	}

	world, err := world.New(fileName, alienCount, alienMaxSteps)
	if err != nil {
		return
	}

	world.Random = random.Random{}
	world.RunAlienInvasion()
}
