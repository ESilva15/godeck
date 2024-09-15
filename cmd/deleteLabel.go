/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func deleteLabelCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	labelId, _ := cmd.Flags().GetInt("label")

	err := deck.DeleteLabel(boardId, labelId)
	if err != nil {
		fmt.Println("Failed to delete label: ", err)
	}
}

// deleteLabelCmd represents the deleteLabel command
var deleteLabelCmd = &cobra.Command{
	Use:   "deleteLabel",
	Short: "",
	Long:  ``,
	Run:   deleteLabelCmdAction,
}

func init() {
	rootCmd.AddCommand(deleteLabelCmd)

	// Set the flags
	deleteLabelCmd.Flags().IntP("board", "b", 0, "Board ID")
	deleteLabelCmd.Flags().IntP("label", "l", 0, "Label ID")

	// Set the required flags
	deleteLabelCmd.MarkFlagRequired("board")
	deleteLabelCmd.MarkFlagRequired("label")
}
