/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func updateBoardCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")

	title, _ := cmd.Flags().GetString("title")
	color, _ := cmd.Flags().GetString("color")
	archived, _ := cmd.Flags().GetBool("archived")

	payload := map[string]interface{}{
		"title":    title,
		"color":    color,
		"archived": archived,
	}

	r, err := deck.UpdateBoard(boardId, payload)
	if err != nil {
		fmt.Println("Could not update card: ", err)
    return
	}

	fmt.Println(r)
}

// updateBoardCmd represents the updateBoard command
var updateBoardCmd = &cobra.Command{
	Use:   "updateBoard",
	Short: "",
	Long:  ``,
	Run:   updateBoardCmdAction,
}

func init() {
	rootCmd.AddCommand(updateBoardCmd)

	updateBoardCmd.Flags().IntP("board", "b", -1, "Board in which to insert the new card")

	updateBoardCmd.Flags().StringP("title", "t", "", "New card title (required)")
	updateBoardCmd.Flags().StringP("color", "c", "FFFFFF", "New board color (default=FFFFFF)")
	updateBoardCmd.Flags().BoolP("archived", "a", false, "Archived state")

	// This ought to be temporary since I have to find a better way of doing it
	updateBoardCmd.MarkFlagRequired("title")

	updateBoardCmd.MarkFlagRequired("boardid")
}
