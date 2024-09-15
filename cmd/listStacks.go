/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func listStacksCmdAction(cmd *cobra.Command, args []string) {
	id, _ := strconv.Atoi(args[0])
	data, err := deck.GetStacks(id)
	if err != nil {
		fmt.Println("Error getting stacks: ", err)
		return
	}

	listCards, _ := cmd.Flags().GetBool("listCards")
  // listArchived, _ := cmd.Flags().GetBool("archived")

	for _, s := range data {
		fmt.Printf("[%-2d] %s\n", s.Id, s.Title)

		if listCards {
			for _, c := range s.Cards {
				fmt.Printf("  [%-2d - %d] %s\n", c.Id, c.Order, c.Title)
			}

      // if listArchived {
      //   deck.GetArchivedCards(boardId, stackId)
      // }
		}
	}
}

// listStacksCmd represents the listStacks command
var listStacksCmd = &cobra.Command{
	Use:   "listStacks [boardId]",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run:   listStacksCmdAction,
}

func init() {
	rootCmd.AddCommand(listStacksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listStacksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listStacksCmd.Flags().BoolP("listCards", "l", false, "Also list cards")
	listStacksCmd.Flags().BoolP("archived", "a", false, "Lists archived cards")
}
