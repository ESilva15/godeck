/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

  "godeck/deck"

	"github.com/spf13/cobra"
)

func assignLabelCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")
  stackId, _ := cmd.Flags().GetInt("stack")
  cardId, _ := cmd.Flags().GetInt("card")
  labelId, _ := cmd.Flags().GetInt("label")

  err := deck.AssignLabelToCard(boardId, stackId, cardId, labelId)
  if err != nil {
    fmt.Println("Failed to assign label to card: ", err)
  }
}

// assignLabelCmd represents the assignLabel command
var assignLabelCmd = &cobra.Command{
	Use:   "assignLabel",
	Short: "",
	Long:  ``,
	Run:   assignLabelCmdAction,
}

func init() {
	rootCmd.AddCommand(assignLabelCmd)

  // Declare the flags for this command
  assignLabelCmd.Flags().IntP("board", "b", 0, "Board ID")
  assignLabelCmd.Flags().IntP("stack", "s", 0, "Stack ID")
  assignLabelCmd.Flags().IntP("card", "c", 0, "Card ID")
  assignLabelCmd.Flags().IntP("label", "l", 0, "Label ID")

  // Mark the required ones
  assignLabelCmd.MarkFlagRequired("board")
  assignLabelCmd.MarkFlagRequired("stack")
  assignLabelCmd.MarkFlagRequired("card")
  assignLabelCmd.MarkFlagRequired("label")
}
