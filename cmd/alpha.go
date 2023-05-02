/*
* Copyright © 2023 Vinuka Kodituwakku <vinuka.t@pm.me>
 */

package cmd

import (
	"github.com/VinukaThejana/generator/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// alphaCmd represents the alpha command
var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "Generate passowrds with only alphabetical characters",
	Long: `
If this option is provided then the generated password will only
contain completly alphabetical characters
  ex :- abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
  `,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			color.Red("Failed to generate password")
			return
		}

		if length <= 2 {
			color.Red("Password length provided is not sufficient")
			return
		}

		generator := utils.Generator{}
		passwd, err := generator.GenerateAlpha(&length)
		if err != nil {
			color.Red("Failed to generate password")
			return
		}

		color.Green(*passwd)
	},
}

func init() {
	rootCmd.AddCommand(alphaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alphaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alphaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
