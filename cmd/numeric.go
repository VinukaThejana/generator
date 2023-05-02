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
	Short: "Generate a numeric only password",
	Long: `
  Generate a password that only consists of numeric only characters
    ex :- 0123456789
  `,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			color.Red("Failed to generate the password")
			return
		}

		if length <= 2 {
			color.Red("Password length is not sufficient to generate the password")
			return
		}

		generator := utils.Passwords{}
		passwd, err := generator.GenerateNumeric(&length)
		if err != nil {
			color.Red("Failed to generate the password")
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
