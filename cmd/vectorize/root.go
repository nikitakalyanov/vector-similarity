package vectorize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "neon-vectorize",
	Short: "neon-vectorize - a CLI for building vector similarity enabled apps",
	Long: `
One can use neon-vectorize to prepare the data in a Neon project for vector similarity search`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
