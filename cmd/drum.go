/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getDrum() int {
	statusINT := snmpGetINT(drumLeftOID)
	return statusINT
}

// drumCmd represents the drum command
var drumCmd = &cobra.Command{
	Use:   "drum",
	Short: "Drum unit life left",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		toner := getDrum()
		fmt.Printf("%d\n", toner)
	},
}

func init() {
	drumCmd.Flags().StringVarP(&host, "hostname", "H", "10.2.0.2", "host to attach")
	drumCmd.Flags().StringVarP(&comm, "community", "c", "v2cpublic", "Community name")
	rootCmd.AddCommand(drumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// drumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// drumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
