/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

  "godeck/deck"

	"github.com/spf13/cobra"
)

func archiveCardCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	stackId, _ := cmd.Flags().GetInt("stack")
	cardId, _ := cmd.Flags().GetInt("card")

  r, err := deck.ArchiveCard(boardId, stackId, cardId)
  if err != nil {
    fmt.Println("Could not archive card: ", err)
  }

  fmt.Println(r)
}

// archiveCardCmd represents the archiveCard command
var archiveCardCmd = &cobra.Command{
	Use:   "archiveCard",
	Short: "",
	Long: ``,
	Run: archiveCardCmdAction,
}

func init() {
	rootCmd.AddCommand(archiveCardCmd)

	archiveCardCmd.Flags().IntP("board", "b", -1, "Board in which to insert the new card")
	archiveCardCmd.Flags().IntP("stack", "s", -1, "Stack in which to insert the new card")
	archiveCardCmd.Flags().IntP("card", "c", -1, "Card to update")

	archiveCardCmd.MarkFlagRequired("boardid")
	archiveCardCmd.MarkFlagRequired("stackid")
}

