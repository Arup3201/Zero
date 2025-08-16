package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCommand = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"subtract"},
	Short:   "Add 2 numbers",
	Long:    "Carry out subtract operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result := Subtract(args[0], args[1])
		if result != "" {
			fmt.Printf("Subtraction of %s from %s = %s.\n\n", args[1], args[0], result)
		}
	},
}

func init() {
	rootCmd.AddCommand(subCommand)
}
