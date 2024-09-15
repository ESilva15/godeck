/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
)

func listCardsCmdAction(cmd *cobra.Command, args []string) {
}

// listCardsCmd represents the listCards command
var listCardsCmd = &cobra.Command{
	Use:   "listCards",
	Short: "",
	Long:  ``,
	Run:   listCardsCmdAction,
}

func init() {
	rootCmd.AddCommand(listCardsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCardsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
