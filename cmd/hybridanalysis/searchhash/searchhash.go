package searchhash

import (
	"fmt"
	"log"

	"github.com/hsmtkk/hybrid_analysis_go/pkg/apikey"
	"github.com/hsmtkk/hybrid_analysis_go/pkg/searchhash"
	"github.com/spf13/cobra"
)

var SearchHashCommand = &cobra.Command{
	Use:  "searchhash",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hash := args[0]
		apiKey, err := apikey.New().LoadAPIKey()
		if err != nil {
			log.Fatal(err)
		}
		result, err := searchhash.New(apiKey).SearchHash(hash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Threat Level: %d\n", result.ThreatLevel)
		fmt.Printf("Threat Score: %d\n", result.ThreatScore)
		fmt.Printf("Verdict: %s\n", result.Verdict)
	},
}
