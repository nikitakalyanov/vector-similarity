package vectorize

import (
	pkg "github.com/nikitakalyanov/vector-similarity/pkg/vectorize"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds a similar object in a database",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.DoFind(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
