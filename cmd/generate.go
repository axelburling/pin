/*
Copyright © 2023 Axel Burling axel.burling@gmail.com
*/
package cmd

import (
	"fmt"
	"time"

	"math/rand"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/axelburling/pin/utils"
	"github.com/spf13/cobra"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [flags]",
	Short: "Generate a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		bday, _ := cmd.Flags().GetString("birthday")
		sep, _ := cmd.Flags().GetBool("separator")
		gender, _ := cmd.Flags().GetString("gender")
		copy, _ := cmd.Flags().GetBool("copy")

		male, _ := cmd.Flags().GetBool("male")
		female, _ := cmd.Flags().GetBool("female")
		long, _ := cmd.Flags().GetBool("long")

		var day string

		switch len(bday) {
		case 0:
			day = start.Format("060102")
		case 6:
			_, err := time.Parse("060102", bday)

			if err != nil {
				fmt.Println("Invalid birthday")
				return
			}

			day = bday

		case 8:
			t, err := time.Parse("20060102", bday)

			if err != nil {
				fmt.Println("Invalid birthday")
				return
			}

			day = t.Format("060102")

		default:
			fmt.Println("Invalid birthday")

		}

		if long {
			now := start.Year()
			year := day[0:2]
			centuryStr := strconv.Itoa(now - ((now - StringToInt(year)) % 100))
			day = fmt.Sprintf("%s%s", centuryStr[0:2], day)
		}

		var g bool

		if male || female {
			g = female
		} else {
			switch gender {
			case "":
				ran := rand.Intn(2)

				g = ran == 1

			case "male":
				g = false

			case "female":
				g = true

			default:
				fmt.Print("Invalid gender supplied")
			}
		}

		val, err := utils.Generate(day, g, sep)

		if err != nil {
			fmt.Println(err)
		}

		if copy {
			clipboard.WriteAll(val)
			fmt.Printf("%s copied to clipboard", val)
		} else {
			fmt.Println(val)
		}

		elapsed := time.Since(start)

		fmt.Printf("Done in %s\n", elapsed)
	},
}

var generatedAliasCmd = &cobra.Command{
	Use:   "gen [flags]",
	Short: "Generate a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {
		generateCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(generatedAliasCmd)

	generateCmd.Flags().BoolP("separator", "s", false, "Add separator to PIN")
	generateCmd.Flags().BoolP("male", "m", false, "")
	generateCmd.Flags().BoolP("female", "f", false, "")
	generateCmd.Flags().BoolP("long", "l", false, "Output Birthday in YYYYMMDD format")
	generateCmd.Flags().BoolP("copy", "c", false, "")
	generateCmd.Flags().StringP("gender", "g", "", "Gender of the generated pin usage ")
	generateCmd.Flags().StringP("birthday", "b", "", "Birthday in YYYYMMDD/YYMMDD format")

}
