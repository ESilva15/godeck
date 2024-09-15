/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func createLabelCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")
  title, _ := cmd.Flags().GetString("title")
  color, _ := cmd.Flags().GetString("color")
  
  payload := map[string]interface{}{
    "title": title,
    "color": color,
  }

  err := deck.CreateLabel(boardId, payload)
  if err != nil {
    fmt.Println("Failed to create label: ", err)
  }
}

// createLabelCmd represents the createLabel command
var createLabelCmd = &cobra.Command{
	Use:   "createLabel",
	Short: "",
	Long:  ``,
	Run:   createLabelCmdAction,
}

func init() {
	rootCmd.AddCommand(createLabelCmd)

  // Set the flags
	createLabelCmd.Flags().IntP("board", "b", 0, "Board ID")
	createLabelCmd.Flags().StringP("title", "t", "New Label", "New label title")
	createLabelCmd.Flags().StringP("color", "c", "FFFFFF", "New label color")

  // Set the required flags
  createLabelCmd.MarkFlagRequired("board")
  createLabelCmd.MarkFlagRequired("title")
  createLabelCmd.MarkFlagRequired("color")
}
