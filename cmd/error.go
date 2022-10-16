/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getPrinterError() (result string) {
	result = "No Error"
	errInt := snmpGetINT(errorStatusOID)
	switch errInt {
	case -1:
		result = "No Error"
	case -2:
		fatalErr := snmpGetINT(fatalErrorOID)
		result = fmt.Sprintf("Fatal error #%d\n", fatalErr)
	case -3:
		opCov := snmpGetINT(coverOpenOID)
		switch opCov {
		case 1:
			result = "Open"
		}
	case -5:
		powSt := snmpGetINT(powderStatusOID)
		switch powSt {
		case 1:
			result = "No cartridge installed"
		case 2:
			result = "Cartridge mismatch"
		case 3:
			result = "No cartridge left"
		}
	case -9:
		outPaper := snmpGetINT(paperOutOID)
		switch outPaper {
		case 1:
			result = "Multipurpose out of paper"
		case 2:
			result = "Automatic feeder out of paper"
		case 4:
			result = "Lotok 1 out of paper"
		case 8:
			result = "Lotok 2 out of paper"
		}
	case -17:
		papJam := snmpGetINT(paperJamOID)
		switch papJam {
		case 0:
			result = "No error"
		default:
			result = "Paper jam"
		}
	default:
		result = fmt.Sprintf("Unknown Error %d, Please get more info", errInt)
	}
	return result
}

// errorCmd represents the error command
var errorCmd = &cobra.Command{
	Use:   "error",
	Short: "check printer error",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		res := getPrinterError()
		fmt.Println(res)
		if res != "No Error" {
			os.Exit(1)
		}
	},
}

func init() {
	errorCmd.Flags().StringVarP(&host, "hostname", "H", "10.2.0.2", "host to attach")
	errorCmd.Flags().StringVarP(&comm, "community", "c", "v2cpublic", "Community name")
	rootCmd.AddCommand(errorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// errorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// errorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
