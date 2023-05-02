/*
* Copyright © 2023 Vinuka Kodituwakku <vinuka.t@pm.me>
 */

// Package cmd contians all the commands that runs in the
// CLI application
package cmd

import (
	"os"

	"github.com/VinukaThejana/generator/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generator",
	Short: "Generate passwords based on the given requirements",
	Long: `Generate passwords based on the given requirements
  Usage:
    Flags:
      --length
          This flag can be provided with any subcommand this flag will be used to
          configure the default length of the password. If not provided the resulting
          password will be 32 characters long

    Commands:
      alpha
        If this option is provided the password generated will only contain alphabetical
        characters only
          ex :- abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

      numeric:
        If this option is provided then the generated password will only contain numeric
        characters only
          ex :- 0123456789

      symbols:
        If this option is provided then the generated password will only contain special
        symbols only.
          ex :- #$%^&*#@@%^^

      base64:
        If this option is provided the resulting password will be encoded in the base 64
        format
          ex :- PQChcGzWyrbVMCLRhoykGw6U9L+R/+/sgx6JorH1mPY=
  `,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			color.Red("Invalid format provided")
			return
		}

		if length <= 2 {
			color.Red("Length provided was invalid")
			return
		}

		generator := utils.Generator{}
		pass, err := generator.Generate(&length)
		if err != nil {
			return
		}

		color.Green(*pass)
		return
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.generator.yaml)")
	rootCmd.PersistentFlags().IntP("length", "l", 32, "Define the length of the password that should be created")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
