/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"github.com/gosnmp/gosnmp"
	"github.com/spf13/cobra"
)

var (
	host             string
	comm             string
	tonerLeftOID     []string = []string{"1.3.6.1.4.1.40093.1.1.3.13.2"}
	errorStatusOID   []string = []string{"1.3.6.1.4.1.40093.1.1.3.15.8"}
	powderStatusOID  []string = []string{"1.3.6.1.4.1.40093.1.1.3.15.3"}
	coverOpenOID     []string = []string{"1.3.6.1.4.1.40093.1.1.3.15.2"}
	paperOutOID      []string = []string{"1.3.6.1.4.1.40093.1.1.3.15.4"}
	paperJamOID      []string = []string{"1.3.6.1.4.1.40093.1.1.3.15.5"}
	fatalErrorOID    []string = []string{"1.3.6.1.4.1.40093.1.1.3.15.1"}
	printerStatusOID []string = []string{"1.3.6.1.4.1.40093.1.1.3.9"}
	drumLeftOID      []string = []string{"1.3.6.1.4.1.40093.8.1.4"}
)

func snmpGetINT(OID []string) int {
	sn := gosnmp.Default
	sn.Target = host
	sn.Community = comm
	err := sn.Connect()
	if err != nil {
		log.Fatalln("Error connnecting to " + host + " via snmp")
	}
	defer sn.Conn.Close()
	res, err := sn.Get(OID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	statusINT := res.Variables[0].Value.(int)
	return statusINT
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pantum-snmp",
	Short: "",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pantum-snmp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
