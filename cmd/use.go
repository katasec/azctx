/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var ctx string

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Uses another azure context",
	Long:  `Uses another azure context`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctx == "" {
			cmd.Help()
			os.Exit(0)
		}

		useProfile(ctx)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	useCmd.Flags().StringVarP(&ctx, "ctx", "c", profile, "Name of context azure context to use")
}

// createProfile
func useProfile(ctx string) {
	fmt.Printf("Switching to '%s' context. \n", ctx)
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
	targetCtx := filepath.Join(azhome, ctx)

	err = CreateSymlink(targetCtx, defaultAzureProfileDir)
	if err != nil {
		fmt.Println("Error creating symlink:", err)
	} else {
		fmt.Println("Symlink created successfully")
	}

}

// CreateSymlink creates a symbolic link pointing to the target folder.
// It checks if the target folder exists before creating the symlink.
func CreateSymlink(target, linkName string) error {
	// Check if the target folder exists
	if _, err := os.Stat(target); os.IsNotExist(err) {
		return fmt.Errorf("target folder does not exist")
	}

	// Remove the symlink if it already exists
	if _, err := os.Lstat(linkName); err == nil {
		if err := os.Remove(linkName); err != nil {
			return fmt.Errorf("failed to remove existing symlink: %w", err)
		}
	}

	// Create the symlink
	err := os.Symlink(target, linkName)
	if err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	return nil
}
