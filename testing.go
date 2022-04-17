/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"


	"github.com/spf13/cobra"
)

// testingCmd represents the testing command
var testingCmd = &cobra.Command{
	Use:   "testing",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}





func init() {
	rootCmd.AddCommand(testingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func getRandomJoke() {
	fmt.Println("Get random dad joke")
    cmd, err := exec.Command("python", "F:/my-calc/testing.py").Output()

if err != nil {
   log.Fatal(err)
 }
 fmt.Println(string(cmd))
 fmt.Println("Done")


  }