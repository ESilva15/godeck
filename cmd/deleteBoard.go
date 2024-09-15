/*
  Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"strconv"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func deleteBoardCmdAction(cmd *cobra.Command, args []string) {
	ids := make([]int, len(args) - 1)
	for _, v := range args {
		id, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Could not read id [%s]: %s\n", v, err)
			continue
		}
		ids = append(ids, id)
	}

	failed := deck.DeleteBoards(ids)
	if len(failed) == 0 {
    fmt.Println("Boards deleted successfully.")
    return
	}

  // Show the user the errors that did happen
  for k, v := range(failed) {
    fmt.Printf("%-3d - %s\n", k, v)
  }
}

// deleteBoardCmd represents the deleteBoard command
var deleteBoardCmd = &cobra.Command{
	Use:   "deleteBoard [IDs]",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run:   deleteBoardCmdAction,
}

func init() {
	rootCmd.AddCommand(deleteBoardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteBoardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteBoardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
