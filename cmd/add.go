package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add 2 numbers",
	Long:    "Carry out addition operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result := Add(args[0], args[1])
		if result != "" {
			fmt.Printf("Addition of %s and %s = %s.\n\n", args[0], args[1], result)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
