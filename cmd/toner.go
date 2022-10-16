/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ()

func getToner() int {
	statusINT := snmpGetINT(tonerLeftOID)
	return statusINT
}

// tonerCmd represents the toner command
var tonerCmd = &cobra.Command{
	Use:   "toner",
	Short: "print toner remained",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		toner := getToner()
		fmt.Printf("%d\n", toner)
	},
}

func init() {
	tonerCmd.Flags().StringVarP(&host, "hostname", "H", "10.2.0.2", "host to attach")
	tonerCmd.Flags().StringVarP(&comm, "community", "c", "v2cpublic", "Community name")
	rootCmd.AddCommand(tonerCmd)
}
