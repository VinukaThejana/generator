/*
* Copyright © 2023 Vinuka Kodituwakku <vinuka.t@pm.me>
 */

package cmd

import (
	"github.com/VinukaThejana/generator/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Generate a password that is encrypted in the basse64 format",
	Long: `
  Generates a password that is encrypted using the base64 encoding mechanism
    ex :- PQChcGzWyrbVMCLRhoykGw6U9L+R/+/sgx6JorH1mPY=
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
		pass, err := generator.GenerateBase64(&length)
		if err != nil {
			color.Red(errors.FailedToGenerate())
			return
		}

		color.Green(*pass)
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
