/*
Copyright © 2022 Kshitij Dhakal dhakalkshitij@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Kshitij-Dhakal/inaugurator/pkg/create"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new boilderplate.",
	Long: `Provide a name for the boilerplate you want to create. Add file for the boilderplate with -f flag.
	
	Example usage : inaugurator create server -f server.go`,
	Run: func(cmd *cobra.Command, args []string) {
		s := &create.BoilderplaterCreatorServiceImpl{}
		f := &create.BoilerplateCreatorImpl{
			Service: s,
		}
		err := f.CreateBoilerplate(file, args...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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
