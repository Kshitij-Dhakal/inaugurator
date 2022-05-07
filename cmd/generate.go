/*
Copyright © 2022 Kshitij Dhakal dhakalkshitij@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Kshitij-Dhakal/inaugurator/facade"
	"github.com/Kshitij-Dhakal/inaugurator/service"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code using the boilderplate you just created.",
	Long: `Generate code using the boilderplate you just created.`,
	Run: func(cmd *cobra.Command, args []string) {
		codeGeneratorService := &service.CodeGeneratorImpl{}
		codeGeneratorFacade := &facade.CodeGeneratorImpl{
			Service: codeGeneratorService,
		}
		err := codeGeneratorFacade.GenerateCode(file, args...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.Flags().StringVarP(&file, "file", "f", "", "file to read")
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
