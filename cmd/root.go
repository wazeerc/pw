/*
Copyright © 2024 wazeerc
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pw",
	Short: "🔐 Effortlessly Generate Robust Passwords!",
	Long: `pw is a minimal CLI application designed to effortlessly generate robust and secure passwords.
You can customize the length and character composition to suit your security needs.`,
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
	// TODO: Run in interactive mode
	// rootCmd.Flags().BoolP("interactive", "i", false, "Run in interactive mode")
}
