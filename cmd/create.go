/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	copy "github.com/katasec/azctx/copy"
	"github.com/spf13/cobra"
)

var profile string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		profile := args[0]
		createProfile(profile)

	},
}

func init() {

	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// createCmd.Flags().StringVarP(&profile, "context", "c", profile, "Name of context to create from current azure settings")
}

// createProfile
func createProfile(ctx string) {

	// Setup homedir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory: " + err.Error())
	}
	azhome := filepath.Join(homeDir, ".azctx")

	// Default Azure profiles stored here
	defaultAzureProfileDir := filepath.Join(homeDir, ".azure")
	fmt.Println("Default azure profile dir:" + defaultAzureProfileDir)

	// New profile will be here
	newProfileDir := filepath.Join(azhome, ctx)
	fmt.Println("New profile dir:" + newProfileDir)

	fmt.Printf("azctx will now copy everything from the default az profile %s to ~/.azctx/%s \n", defaultAzureProfileDir, newProfileDir)

	copy.CopyDir(defaultAzureProfileDir, newProfileDir)
}
