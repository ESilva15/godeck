/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func viewBoardCmdAction(cmd *cobra.Command, args []string) {
  boardId, _ := cmd.Flags().GetInt("board")

	board, err := deck.GetBoardById(boardId)
	if err != nil {
		fmt.Println("Could not retrieve boards: ", err)
    return
	}

  board.Show()
}

// viewBoardCmd represents the viewBoard command
var viewBoardCmd = &cobra.Command{
	Use:   "viewBoard",
	Short: "",
	Long:  ``,
	Run:   viewBoardCmdAction,
}

func init() {
	rootCmd.AddCommand(viewBoardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewBoardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	viewBoardCmd.Flags().IntP("board", "b", 0, "Board ID")
}
