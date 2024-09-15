/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "godeck/deck"

	"github.com/spf13/cobra"
)

func reorderCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")
  stackId, _ := cmd.Flags().GetInt("stack")
  cardId, _ := cmd.Flags().GetInt("card")
  order, _ := cmd.Flags().GetInt("newOrder")
  stack, _ := cmd.Flags().GetInt("newStack")

  deck.ReorderCard(boardId, stackId, cardId, order, stack)
}

// reorderCmd represents the reorder command
var reorderCmd = &cobra.Command{
	Use:   "reorder",
	Short: "",
	Long:  ``,
	Run:   reorderCmdAction,
}

func init() {
	rootCmd.AddCommand(reorderCmd)

  // Set the flags
	reorderCmd.Flags().IntP("board", "b", 0, "Board ID")
	reorderCmd.Flags().IntP("stack", "s", 0, "Stack ID")
	reorderCmd.Flags().IntP("card", "c", 0, "Card ID")
	reorderCmd.Flags().IntP("newOrder", "o", 0, "Card position in stack")
	reorderCmd.Flags().IntP("newStack", "n", 0, "New stack")

  // Mark the required flags
  reorderCmd.MarkFlagRequired("board")
  reorderCmd.MarkFlagRequired("stack")
  reorderCmd.MarkFlagRequired("card")
  reorderCmd.MarkFlagRequired("order")
  reorderCmd.MarkFlagRequired("newStack")
}
