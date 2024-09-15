/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"

  "godeck/deck"
  
	"github.com/spf13/cobra"
)

var (
  title string
  color string
)

// createBoardCmd represents the createBoard command
var createBoardCmd = &cobra.Command{
	Use:   "createBoard",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
    deck.CreateBoard(title, color)
	},
}

func init() {
	rootCmd.AddCommand(createBoardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createBoardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createBoardCmd.Flags().StringVarP(&title, "title", "t",
		"New Board", "New board title\nMax chars: XXX.",
	)
	createBoardCmd.Flags().StringVarP(&color, "color", "c",
		"000000", "New board color.",
	)

  createBoardCmd.MarkFlagRequired("title")
  createBoardCmd.MarkFlagRequired("color")
}
