/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "fmt"

  "godeck/deck"

	"github.com/spf13/cobra"
)

func removeLabelCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")
  stackId, _ := cmd.Flags().GetInt("stack")
  cardId, _ := cmd.Flags().GetInt("card")
  labelId, _ := cmd.Flags().GetInt("label")

  err := deck.UnassignLabelFromCard(boardId, stackId, cardId, labelId)
  if err != nil {
    fmt.Println("Failed to assign label to card: ", err)
  }
}

// removeLabelCmd represents the removeLabel command
var removeLabelCmd = &cobra.Command{
	Use:   "removeLabel",
	Short: "",
	Long:  ``,
	Run:   removeLabelCmdAction,
}

func init() {
	rootCmd.AddCommand(removeLabelCmd)

  // Declare the flags for this command
  removeLabelCmd.Flags().IntP("board", "b", 0, "Board ID")
  removeLabelCmd.Flags().IntP("stack", "s", 0, "Stack ID")
  removeLabelCmd.Flags().IntP("card", "c", 0, "Card ID")
  removeLabelCmd.Flags().IntP("label", "l", 0, "Label ID")

  // Mark the required ones
  removeLabelCmd.MarkFlagRequired("board")
  removeLabelCmd.MarkFlagRequired("stack")
  removeLabelCmd.MarkFlagRequired("card")
  removeLabelCmd.MarkFlagRequired("label")
}
