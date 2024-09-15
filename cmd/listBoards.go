/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func listBoardsCmdAction(cmd *cobra.Command, args []string) {
	boards, err := deck.GetBoards()
	if err != nil {
		fmt.Println("Could not retrieve boards: ", err)
	}

	archived, _ := cmd.Flags().GetBool("archived")
	deleted, _ := cmd.Flags().GetBool("deleted")

	for k, v := range boards {
		if archived && v.Archived {
			fmt.Printf("A [%-2d] %s\n", k, v.Title)
		} else if deleted && v.DeletedAt > 0 {
			fmt.Printf("D [%-2d] %s\n", k, v.Title)
		} else if v.DeletedAt <= 0 && !v.Archived {
      fmt.Printf("  [%-2d] %s\n", k, v.Title)
    }
	}
}

// listBoardsCmd represents the listBoards command
var listBoardsCmd = &cobra.Command{
	Use:   "listBoards [IDs]",
	Short: "Lists the Deck boards.",
	Long:  ``,
	Run:   listBoardsCmdAction,
}

func init() {
	rootCmd.AddCommand(listBoardsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listBoardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listBoardsCmd.Flags().BoolP("archived", "a", false, "Shows archived boards")
	listBoardsCmd.Flags().BoolP("deleted", "d", false, "Shows deleted boards")
}
