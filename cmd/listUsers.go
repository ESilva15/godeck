/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func listUsersCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")

  users, err := deck.GetUsers(boardId)
  if err != nil {
    fmt.Println("Could not get users: ", err)
    return
  }

  for _, u := range(users) {
    u.Show()
  }
}

// listUsersCmd represents the listUsers command
var listUsersCmd = &cobra.Command{
	Use:   "listUsers",
	Short: "",
	Long:  ``,
	Run:   listUsersCmdAction,
}

func init() {
	rootCmd.AddCommand(listUsersCmd)

  // Set command flags
  listUsersCmd.Flags().IntP("board", "b", 0, "Board ID")

  // Mark required flags
  listUsersCmd.MarkFlagRequired("board")
}
