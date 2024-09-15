/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func updateLabelCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	labelId, _ := cmd.Flags().GetInt("label")
	title, _ := cmd.Flags().GetString("title")
	color, _ := cmd.Flags().GetString("color")

	payload := map[string]interface{}{
		"title": title,
		"color": color,
	}

	err := deck.UpdateLabel(boardId, labelId, payload)
	if err != nil {
		fmt.Println("Failed to update label: ", err)
	}
}

// updateLabelCmd represents the updateLabel command
var updateLabelCmd = &cobra.Command{
	Use:   "updateLabel",
	Short: "",
	Long:  ``,
	Run:   updateLabelCmdAction,
}

func init() {
	rootCmd.AddCommand(updateLabelCmd)

	// Set the flags
	updateLabelCmd.Flags().IntP("board", "b", 0, "Board ID")
	updateLabelCmd.Flags().IntP("label", "l", 0, "Label ID")
	updateLabelCmd.Flags().StringP("title", "t", "New Label", "New label title")
	updateLabelCmd.Flags().StringP("color", "c", "FFFFFF", "New label color")

	// Set the required flags
	updateLabelCmd.MarkFlagRequired("board")
	updateLabelCmd.MarkFlagRequired("label")
	updateLabelCmd.MarkFlagRequired("title")
	updateLabelCmd.MarkFlagRequired("color")
}

