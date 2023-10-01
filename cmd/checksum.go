/*
Copyright © 2023 Axel Burling axel.burling@gmail.com
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/axelburling/pin/utils"
	"github.com/spf13/cobra"
)

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:     "checksum <pin> | pin cs <pin>",
	Aliases: []string{"cs"},
	Short:   "Calculate the checksum for a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()

		copy, _ := cmd.Flags().GetBool("copy")

		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}

		pin := args[0]

		p := utils.Pin{}

		parsedPin, err := p.ParseCheck(pin)

		if err != nil {
			fmt.Println("Invalid PIN")
			return
		}

		v := p.VerifyTime(parsedPin[0:6])

		if !v {
			fmt.Println("Invalid PIN")
			return
		}

		if err != nil {
			fmt.Println("Invalid PIN")
			return
		}

		l := utils.Luhn{Pin: parsedPin}

		val, err := l.Calculate()

		if err != nil {
			fmt.Println("Invalid PIN")
			return
		}

		if copy {
			clipboard.WriteAll(fmt.Sprintf("%s%s", parsedPin, val))

			fmt.Printf("Copied to clipboard")
		} else {
			fmt.Printf("The checksum for %s is %s\n", parsedPin, val)
		}

		elapsed := time.Since(start)

		fmt.Printf("Done in %s\n", elapsed)

	},
}

var checksumAliasCmd = &cobra.Command{
	Use:   "cs [flags]",
	Short: "Calculate the checksum for a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {
		checksumCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)
	rootCmd.AddCommand(checksumAliasCmd)

	checksumCmd.Flags().BoolP("copy", "c", false, "")
}
