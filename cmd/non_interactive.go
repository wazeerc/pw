package cmd

import (
	utils "pw/utils"

	"github.com/spf13/cobra"
)

// This is the default command that runs in non-interactive mode, i.e. it uses flags.
var nonInteractiveMode = &cobra.Command{
	Use:   "gen",
	Short: "Generate a password",
	Long: `Generate a password with the specified length and character composition.
For example, to generate a password with a length of 16 characters and 4 digits, 
you would run the following command: pw generate -l 16 -n -s`,

	Run: getPassword,
}

func init() {
	rootCmd.AddCommand(nonInteractiveMode)

	nonInteractiveMode.Flags().IntP("length", "l", 12, "Length of the password")
	nonInteractiveMode.Flags().BoolP("numbers", "n", false, "Include digits in the password")
	nonInteractiveMode.Flags().BoolP("symbols", "s", false, "Include symbols in the password")
}

func getPassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("numbers")
	isSymbols, _ := cmd.Flags().GetBool("symbols")

	password := utils.GeneratePassword(length, isDigits, isSymbols)

	cmd.Println("📋 Your password has been copied to your clipboard!")
	utils.WriteToClipboard(password)
}
