/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func listLabelsCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")

	labels, err := deck.GetLabels(boardId)
	if err != nil {
		fmt.Println("Could not retrieve boards: ", err)
	}

	for _, l := range labels {
		l.Show()
	}
}

// listLabelsCmd represents the listLabels command
var listLabelsCmd = &cobra.Command{
	Use:   "listLabels",
	Short: "",
	Long:  ``,
	Run:   listLabelsCmdAction,
}

func init() {
	rootCmd.AddCommand(listLabelsCmd)

	listLabelsCmd.Flags().IntP("board", "b", 0, "ID of the board")

	listLabelsCmd.MarkFlagRequired("board")
}
