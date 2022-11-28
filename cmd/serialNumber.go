/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getSerialNumber() string {
	res := string(snmpGetASCII(serialNumberOID))
	return res
}

// serialNumberCmd represents the serialNumber command
var serialNumberCmd = &cobra.Command{
	Use:   "serialNumber",
	Short: "print serial number",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getSerialNumber())
	},
}

func init() {
	serialNumberCmd.Flags().StringVarP(&host, "hostname", "H", "10.2.0.2", "host to attach")
	serialNumberCmd.Flags().StringVarP(&comm, "community", "c", "v2cpublic", "Community name")
	rootCmd.AddCommand(serialNumberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serialNumberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serialNumberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
