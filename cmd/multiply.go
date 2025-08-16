package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var mulCommand = &cobra.Command{
	Use:     "mul",
	Aliases: []string{"multi", "multiple", "multiply"},
	Short:   "Multiply 2 numbers",
	Long:    "Carry out multiplication operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result := Multiply(args[0], args[1], shouldRoundUp)
		if result != "" {
			fmt.Printf("Multiplicatoin of %s and %s = %s.\n\n", args[0], args[1], result)
		}
	},
}

func init() {
	mulCommand.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
	rootCmd.AddCommand(mulCommand)
}
