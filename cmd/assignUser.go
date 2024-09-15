/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "godeck/deck"

	"github.com/spf13/cobra"
)

func assignUserCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")
  stackId, _ := cmd.Flags().GetInt("stack")
  cardId, _ := cmd.Flags().GetInt("card")
  userId, _ := cmd.Flags().GetStringSlice("user")

  deck.AssignUserToCard(boardId, stackId, cardId, userId)
}

// assignUserCmd represents the assignUser command
var assignUserCmd = &cobra.Command{
	Use:   "assignUser",
	Short: "",
	Long: ``,
	Run: assignUserCmdAction,
}

func init() {
	rootCmd.AddCommand(assignUserCmd)

  // Declare the flags for this command
  assignUserCmd.Flags().IntP("board", "b", 0, "Board ID")
  assignUserCmd.Flags().IntP("stack", "s", 0, "Stack ID")
  assignUserCmd.Flags().IntP("card", "c", 0, "Card ID")
  assignUserCmd.Flags().StringSliceP("user", "u", []string{}, "User ID")

  // Mark the required ones
  assignUserCmd.MarkFlagRequired("board")
  assignUserCmd.MarkFlagRequired("stack")
  assignUserCmd.MarkFlagRequired("card")
  assignUserCmd.MarkFlagRequired("user")
}
