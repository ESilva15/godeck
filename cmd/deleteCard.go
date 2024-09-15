/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func deleteCardCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	stackId, _ := cmd.Flags().GetInt("stack")
	cardIds, _ := cmd.Flags().GetIntSlice("cards")

  for _, c := range(cardIds) {
    r, err := deck.DeleteCard(boardId, stackId, c)
    if err != nil {
      fmt.Println("Unable to delete card: ", err)
      continue
    }

    fmt.Println(r)
  }
}

// deleteCardCmd represents the deleteCard command
var deleteCardCmd = &cobra.Command{
	Use:   "deleteCard",
	Short: "",
	Long: ``,
	Run: deleteCardCmdAction,
}

func init() {
	rootCmd.AddCommand(deleteCardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCardCmd.Flags().IntP("board", "b", -1, "Board in which to insert the new card")
	deleteCardCmd.Flags().IntP("stack", "s", -1, "Stack in which to insert the new card")
	deleteCardCmd.Flags().IntSliceP("cards", "c", []int{}, "Card or list of cards to delete")

	deleteCardCmd.MarkFlagRequired("board")
	deleteCardCmd.MarkFlagRequired("stack")
	deleteCardCmd.MarkFlagRequired("cards")
}
