/*
Copyright © 2023 Axel Burling axel.burling@gmail.com
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/axelburling/pin/utils"
	"github.com/spf13/cobra"
)

var (
	monthMap = map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain <pin>",
	Short: "Explain the structure of a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		checksum, _ := cmd.Flags().GetBool("checksum")

		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}

		pin := args[0]

		p := utils.Pin{}

		v := p.Verify(pin)

		if !v {
			fmt.Println("Invalid PIN")
			return
		}

		var gender string

		if utils.StringToInt(p.Gender)%2 == 0 {
			gender = "Female"
		} else {
			gender = "Male"
		}

		place := p.GetPlaceOfBirth()

		month := monthMap[p.Month]

		fmt.Printf("Birthday: %s %s, %s\n", p.Day, month, p.FullYear)
		fmt.Printf("Gender: %s\n", gender)
		fmt.Printf("Place of birth: %s\n", place)
		if checksum {
			fmt.Printf("Control Digit: %s\n", p.ControlDigit)
		}

		elapsed := time.Since(start)

		fmt.Printf("Done in %s\n", elapsed)
	},
}

var explainAliasCmd = &cobra.Command{
	Use:   "ex <pin>",
	Short: "Explain the structure of a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {
		explainCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
	rootCmd.AddCommand(explainAliasCmd)

	explainCmd.Flags().BoolP("checksum", "c", false, "Show checksum(It is sensitive information, should not be shared with others)")
}
