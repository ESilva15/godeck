/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func createStackCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	order, _ := cmd.Flags().GetInt("order")
	title, _ := cmd.Flags().GetString("title")

	payload := map[string]interface{}{
		"order": order,
		"title": title,
	}

  r, err := deck.CreateStack(boardId, payload)
  if err != nil {
    fmt.Println("Failed to create stack: ", err)
    return
  }

  fmt.Println(r)
}

// createStackCmd represents the createStack command
var createStackCmd = &cobra.Command{
	Use:   "createStack",
	Short: "",
	Long:  ``,
	Run:   createStackCmdAction,
}

func init() {
	rootCmd.AddCommand(createStackCmd)

	createStackCmd.Flags().IntP("board", "b", 0, "Board ID")
	createStackCmd.Flags().IntP("order", "o", 0, "Order of the stack")
	createStackCmd.Flags().StringP("title", "t", "New Title", "New title for the stack")

	createStackCmd.MarkFlagRequired("board")
	createStackCmd.MarkFlagRequired("order")
	createStackCmd.MarkFlagRequired("title")
}
