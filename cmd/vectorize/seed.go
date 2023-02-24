package vectorize

import (
	pkg "github.com/nikitakalyanov/vector-similarity/pkg/vectorize"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeds a database",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.DoSeed(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
