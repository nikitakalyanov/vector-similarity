package vectorize

import (
	pkg "github.com/nikitakalyanov/vector-similarity/pkg/vectorize"
	"github.com/spf13/cobra"
)

var vectorizeCmd = &cobra.Command{
	Use:   "vectorize",
	Short: "Vectorizes records in the database",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.DoVectorize(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(vectorizeCmd)
}
