/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getPrintStatus() (status string) {
	statusINT := snmpGetINT(printerStatusOID)
	switch statusINT {
	case 0:
		status = "Init"
	case 1:
		status = "Sleep"
	case 2:
		status = "Preheat"
	case 3:
		status = "Standby"
	case 4:
		status = "Printing"
	case 5:
		status = "Error"
	case 6:
		status = "Cancel the job"
	case 7:
		status = "Proccessing"
	default:
		status = "Unknown status"
	}
	return
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Current state of pantum",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		status := getPrintStatus()
		fmt.Println(status)
		if status == "Error" {
			os.Exit(1)
		}
	},
}

func init() {

	statusCmd.Flags().StringVarP(&host, "hostname", "H", "10.2.0.2", "host to attach")
	statusCmd.Flags().StringVarP(&comm, "community", "c", "v2cpublic", "Community name")
	rootCmd.AddCommand(statusCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
