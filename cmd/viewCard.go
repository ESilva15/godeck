/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func viewCardCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	stackId, _ := cmd.Flags().GetInt("stack")
	cardIds, _ := cmd.Flags().GetIntSlice("card")

	cards, err := deck.GetCards(boardId, stackId, cardIds)
	if err != nil {
		fmt.Println("Unable to get cards: ", err)
	}

	for _, c := range cards {
		c.Show()
	}
}

// viewCardCmd represents the viewCard command
var viewCardCmd = &cobra.Command{
	Use:   "viewCard",
	Short: "",
	Long:  ``,
	// Args:  cobra.MinimumNArgs(1),
	Run: viewCardCmdAction,
}

func init() {
	rootCmd.AddCommand(viewCardCmd)

  // Set the flags
	viewCardCmd.Flags().IntP("board", "b", 0, "ID of the board to lookup.")
	viewCardCmd.Flags().IntP("stack", "s", 0, "ID of the stack to lookup.")
	viewCardCmd.Flags().IntSliceP("card", "c", []int{}, "IDs of the cards to lookup.")

  // Mark required falgs
  viewCardCmd.MarkFlagRequired("board")
  viewCardCmd.MarkFlagRequired("stack")
  viewCardCmd.MarkFlagRequired("card")
}
