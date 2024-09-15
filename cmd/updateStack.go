/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func updateStackCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	stackId, _ := cmd.Flags().GetInt("stack")
	order, _ := cmd.Flags().GetInt("order")
	title, _ := cmd.Flags().GetString("title")

	payload := map[string]interface{}{
		"order": order,
		"title": title,
	}

  r, err := deck.UpdateStack(boardId, stackId, payload)
  if err != nil {
    fmt.Println("Failed to create stack: ", err)
    return
  }

  fmt.Println(r)
}

// updateStackCmd represents the updateStack command
var updateStackCmd = &cobra.Command{
	Use:   "updateStack",
	Short: "",
	Long:  ``,
	Run:   updateStackCmdAction,
}

func init() {
	rootCmd.AddCommand(updateStackCmd)

	updateStackCmd.Flags().IntP("board", "b", 0, "Board ID")
	updateStackCmd.Flags().IntP("stack", "s", 0, "Stack ID")
	updateStackCmd.Flags().IntP("order", "o", 0, "Order of the stack")
	updateStackCmd.Flags().StringP("title", "t", "New Title", "New title for the stack")

	updateStackCmd.MarkFlagRequired("board")
	updateStackCmd.MarkFlagRequired("stack")
	updateStackCmd.MarkFlagRequired("order")
	updateStackCmd.MarkFlagRequired("title")
}
