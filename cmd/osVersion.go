/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getOSVersion() string {
	res := string(snmpGetASCII(firmwareOID))
	return res
}

// osVersionCmd represents the osVersion command
var osVersionCmd = &cobra.Command{
	Use:   "osVersion",
	Short: "print of version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		res := getOSVersion()
		fmt.Println(res)
	},
}

func init() {
	osVersionCmd.Flags().StringVarP(&host, "hostname", "H", "10.2.0.2", "host to attach")
	osVersionCmd.Flags().StringVarP(&comm, "community", "c", "v2cpublic", "Community name")
	rootCmd.AddCommand(osVersionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// osVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// osVersionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
