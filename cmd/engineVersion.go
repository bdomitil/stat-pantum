/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getEngineVersion() string {
	res := string(snmpGetASCII(engineVersionOID))
	return res
}

// engineVersionCmd represents the engineVersion command
var engineVersionCmd = &cobra.Command{
	Use:   "engineVersion",
	Short: "print engine version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getEngineVersion())
	},
}

func init() {
	engineVersionCmd.Flags().StringVarP(&host, "hostname", "H", "10.2.0.2", "host to attach")
	engineVersionCmd.Flags().StringVarP(&comm, "community", "c", "v2cpublic", "Community name")
	rootCmd.AddCommand(engineVersionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// engineVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// engineVersionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
