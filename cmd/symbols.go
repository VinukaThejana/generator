/*
* Copyright © 2023 Vinuka Kodituwakku <vinuka.t@pm.me>
 */

package cmd

import (
	"github.com/VinukaThejana/generator/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// symbolsCmd represents the symbols command
var symbolsCmd = &cobra.Command{
	Use:   "symbols",
	Short: "Generate a password that consists only symbols",
	Long: `
Generate a password that only consists of symbols
  ex :- #$%^&*#@@%^^
`,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			color.Red("Failed to generate password")
			return
		}

		if length <= 2 {
			color.Red("Failed to generate password")
			return
		}

		generator := utils.Generator{}
		paswd, err := generator.GenerateSpecialChars(&length)
		if err != nil {
			color.Red("Failed to generate password")
			return
		}

		color.Green(*paswd)
	},
}

func init() {
	rootCmd.AddCommand(symbolsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// symbolsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// symbolsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
