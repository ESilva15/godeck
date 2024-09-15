/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

  "godeck/deck"

	"github.com/spf13/cobra"
)

func updateCardCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	stackId, _ := cmd.Flags().GetInt("stack")
	cardId, _ := cmd.Flags().GetInt("card")

  title, _ := cmd.Flags().GetString("title")
  cType, _ := cmd.Flags().GetString("type")
  order, _ := cmd.Flags().GetInt("order")
  description, _ := cmd.Flags().GetString("description")
  duedate, _ := cmd.Flags().GetString("duedate")
  owner, _ := cmd.Flags().GetString("owner")
  payload := map[string]interface{} {
    "title": title,
    "type": cType,
    "order": order,
    "description": description,
    "duedate": duedate,
    "owner": owner,
  }

  fmt.Println("Payload: ")
  fmt.Println(payload)

  r, err := deck.UpdateCard(boardId, stackId, cardId, payload)
  if err != nil {
    fmt.Println("Could not update card: ", err)
  }

  fmt.Println(r)
}

// updateCardCmd represents the updateCard command
var updateCardCmd = &cobra.Command{
	Use:   "updateCard",
	Short: "",
	Long: ``,
	Run: updateCardCmdAction,
}

func init() {
	rootCmd.AddCommand(updateCardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	updateCardCmd.Flags().IntP("board", "b", -1, "Board in which to insert the new card")
	updateCardCmd.Flags().IntP("stack", "s", -1, "Stack in which to insert the new card")
	updateCardCmd.Flags().IntP("card", "c", -1, "Card to update")

	updateCardCmd.Flags().StringP("title", "t", "", "New card's title (required)")
	updateCardCmd.Flags().StringP("type", "y", "plain", "New card's type (default=plain)")
	updateCardCmd.Flags().IntP("order", "o", 1, "New card's order (default=1)")
	updateCardCmd.Flags().StringP("description", "e", "", "New card's description (optional)")
	updateCardCmd.Flags().StringP("duedate", "d", "", "New card's duedate (optional)")
	updateCardCmd.Flags().StringP("owner", "w", "username", "Cards owner (required)")

  // This ought to be temporary since I have to find a better way of doing it
	updateCardCmd.MarkFlagRequired("title")
	updateCardCmd.MarkFlagRequired("type")
	updateCardCmd.MarkFlagRequired("order")
	updateCardCmd.MarkFlagRequired("description")
	updateCardCmd.MarkFlagRequired("duedate")
	updateCardCmd.MarkFlagRequired("owner")
  //

	updateCardCmd.MarkFlagRequired("boardid")
	updateCardCmd.MarkFlagRequired("stackid")
}
