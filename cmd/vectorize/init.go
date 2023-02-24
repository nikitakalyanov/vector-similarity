package vectorize

import (
	//"fmt"

	pkg "github.com/nikitakalyanov/vector-similarity/pkg/vectorize"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.DoInit(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
