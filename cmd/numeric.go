/*
* Copyright © 2023 Vinuka Kodituwakku <vinuka.t@pm.me>
 */

package cmd

import (
	"github.com/VinukaThejana/generator/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// numericCmd represents the numeric command
var numericCmd = &cobra.Command{
	Use:   "numeric",
	Short: "Generate a password that only contains numeric characters",
	Long: `
  Generates a password that only contains numbers
    ex :- 0123456789
  `,
	Run: func(cmd *cobra.Command, args []string) {
		errors := utils.Errors{}
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			color.Red(errors.FailedToGenerate())
			return
		}

		if length <= 2 {
			color.Red(errors.ToShort())
			return
		}

		generator := utils.Generator{}
		passwd, err := generator.GenerateNumeric(&length)
		if err != nil {
			color.Red(errors.FailedToGenerate())
			return
		}

		color.Green(*passwd)
	},
}

func init() {
	rootCmd.AddCommand(numericCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numericCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numericCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
