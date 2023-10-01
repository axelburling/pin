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

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify <pin>",
	Short: "Verify a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()

		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}

		pin := args[0]

		p := utils.Pin{}
		res := p.Verify(pin)

		if res {
			fmt.Println("PIN is valid")
		} else {
			fmt.Println("PIN is invalid")
		}

		elapsed := time.Since(start)

		fmt.Printf("Done in %s\n", elapsed)
	},
}

var verifyAliasCmd = &cobra.Command{
	Use:   "v <pin>",
	Short: "Verify a PIN(Personal Identification Number)",
	Run: func(cmd *cobra.Command, args []string) {
		verifyCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	rootCmd.AddCommand(verifyAliasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
