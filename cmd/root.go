/*
Copyright © 2023 Axel Burling axel.burling@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pin",
	Short: "Manage PINs(Swedish personal identity number)",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
