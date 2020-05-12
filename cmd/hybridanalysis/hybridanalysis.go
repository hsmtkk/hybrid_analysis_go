package main

import (
	"log"

	"github.com/hsmtkk/hybrid_analysis_go/cmd/hybridanalysis/searchhash"
	"github.com/spf13/cobra"
)

func main() {
	rootCommand := &cobra.Command{
		Use: "hybridanalysis",
	}
	rootCommand.AddCommand(searchhash.SearchHashCommand)
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
