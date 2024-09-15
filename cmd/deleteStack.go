/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func deleteStackCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	stackId, _ := cmd.Flags().GetInt("stack")

	err := deck.DeleteLabel(boardId, stackId)
	if err != nil {
		fmt.Println("Failed to delete label: ", err)
	}
}

// deleteStackCmd represents the deleteStack command
var deleteStackCmd = &cobra.Command{
	Use:   "deleteStack",
	Short: "",
	Long:  ``,
	Run:   deleteStackCmdAction,
}

func init() {
	rootCmd.AddCommand(deleteStackCmd)

	// Set the flags
	deleteStackCmd.Flags().IntP("board", "b", 0, "Board ID")
	deleteStackCmd.Flags().IntP("stack", "s", 0, "Label ID")

	// Set the required flags
	deleteStackCmd.MarkFlagRequired("board")
	deleteStackCmd.MarkFlagRequired("stack")
}

