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
	Short: "Generate a password based on the base64 encoding",
	Long: `
  Generates a password that is encoded in the base64 encoding
    ex :- PQChcGzWyrbVMCLRhoykGw6U9L+R/+/sgx6JorH1mPY=
  `,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			color.Red("Failed to generate password")
			return
		}

		if length <= 2 {
			color.Red("Password length is not sufficent for password generation")
			return
		}

		generator := utils.Passwords{}
		pass, err := generator.GenerateBase64(&length)
		if err != nil {
			color.Red("Failed to generate password")
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
