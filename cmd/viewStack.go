/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"godeck/deck"

	"github.com/spf13/cobra"
)

func viewStackCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	stackId, _ := cmd.Flags().GetInt("stack")

	stack, err := deck.GetStack(boardId, stackId)
	if err != nil {
		fmt.Println("Unable to get stack: ", err)
		return
	}

	stack.Show()
}

// viewStackCmd represents the viewStack command
var viewStackCmd = &cobra.Command{
	Use:   "viewStack",
	Short: "",
	Long:  ``,
	Run:   viewStackCmdAction,
}

func init() {
	rootCmd.AddCommand(viewStackCmd)

	viewStackCmd.Flags().IntP("board", "b", 0, "Board ID")
	viewStackCmd.Flags().IntP("stack", "s", 0, "Stack ID")
}
