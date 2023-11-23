package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1"

var rootCmd = &cobra.Command{
	Use:     "cut_something",
	Version: version,
	Short:   "A little helper for cutting things (e.g. used in woodworking)",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
