/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func createCardCmdAction(cmd *cobra.Command, args []string) {
	nTitle, _ := cmd.Flags().GetString("title")
	nType, _ := cmd.Flags().GetString("type")
	nOrder, _ := cmd.Flags().GetInt("order")
	nDescription, _ := cmd.Flags().GetString("description")
	nDuedate, _ := cmd.Flags().GetInt("duedate")

	boardId, _ := cmd.Flags().GetInt("boardid")
	stackId, _ := cmd.Flags().GetInt("stackid")

	createCard := map[string]interface{}{
		"title": nTitle,
		"type":  nType,
		"order": nOrder,
	}

	if nDuedate != -1 {
		createCard["duedata"] = nDuedate
	}
	if nDescription != "" {
		createCard["description"] = nDescription
	}

  r, err := deck.CreateCard(boardId, stackId, createCard)
  if err != nil {
    fmt.Println("There was an error: ", err)
    return
  }

  fmt.Println(r)
}

// createCardCmd represents the createCard command
var createCardCmd = &cobra.Command{
	Use:   "createCard",
	Short: "",
	Long:  ``,
	Run:   createCardCmdAction,
}

func init() {
	rootCmd.AddCommand(createCardCmd)

	createCardCmd.Flags().IntP("board", "b", -1, "Board in which to insert the new card.")
	createCardCmd.Flags().IntP("stack", "s", -1, "Stack in which to insert the new card.")

	createCardCmd.Flags().StringP("title", "t", "", "New card's title (required)")
	createCardCmd.Flags().StringP("type", "y", "plain", "New card's type (default=plain)")
	createCardCmd.Flags().IntP("order", "o", 1, "New card's order (default=1)")
	createCardCmd.Flags().StringP("description", "d", "", "New card's description (optional)")
	createCardCmd.Flags().IntP("duedate", "c", -1, "New card's duedate (optional, $(date -d \"+1 week\" +%s))")

	createCardCmd.MarkFlagRequired("title")
	createCardCmd.MarkFlagRequired("boardid")
	createCardCmd.MarkFlagRequired("stackid")
}
