/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "godeck/deck"

	"github.com/spf13/cobra"
)

func unassignUserCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")
  stackId, _ := cmd.Flags().GetInt("stack")
  cardId, _ := cmd.Flags().GetInt("card")
  userId, _ := cmd.Flags().GetStringSlice("user")

  deck.UnassignUserFromCard(boardId, stackId, cardId, userId)
}

// unassignUserCmd represents the assignUser command
var unassignUserCmd = &cobra.Command{
	Use:   "unassignUser",
	Short: "",
	Long: ``,
	Run: unassignUserCmdAction,
}

func init() {
	rootCmd.AddCommand(unassignUserCmd)

  // Declare the flags for this command
  unassignUserCmd.Flags().IntP("board", "b", 0, "Board ID")
  unassignUserCmd.Flags().IntP("stack", "s", 0, "Stack ID")
  unassignUserCmd.Flags().IntP("card", "c", 0, "Card ID")
  unassignUserCmd.Flags().StringSliceP("user", "u", []string{}, "User ID")

  // Mark the required ones
  unassignUserCmd.MarkFlagRequired("board")
  unassignUserCmd.MarkFlagRequired("stack")
  unassignUserCmd.MarkFlagRequired("card")
  unassignUserCmd.MarkFlagRequired("user")
}
