/*
Copyright © 2022 Kshitij Dhakal dhakalkshitij@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Kshitij-Dhakal/inaugurator/pkg/boilerplate"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the boilerplates you have created.",
	Long: `List the boilerplates you have created.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := &boilerplate.BoilerplateListerServiceImpl{}
		f := &boilerplate.BoilerplateListerFacadeImpl{
			Service: s,
		}
		err := f.ListBoilerplates()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
