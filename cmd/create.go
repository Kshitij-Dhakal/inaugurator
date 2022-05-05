/*
Copyright © 2022 Kshitij Dhakal dhakalkshitij@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create command",
	Long: `Create command`,
	Run: func(cmd *cobra.Command, args []string) {
		if(file == ""){
			fmt.Println("Please provide a file")
			os.Exit(1)
		}
		fmt.Println("Creating boilerplate from File : ", file)
		
		dat, err := os.ReadFile(file)
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(dat))
	},
}

var file string

func init() {
	createCmd.Flags().StringVarP(&file, "file", "f", "", "file to read")
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
