/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "fmt"

  "godeck/deck"

	"github.com/spf13/cobra"
)

func viewLabelCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
  labelIds, _ := cmd.Flags().GetIntSlice("label")

  for _, lId := range(labelIds) {
    label, err := deck.GetLabel(boardId, lId)
    if err != nil {
      fmt.Println("Could not retrieve boards: ", err)
    }

    label.Show()
  }
}

// viewLabelCmd represents the viewLabel command
var viewLabelCmd = &cobra.Command{
	Use:   "viewLabel",
	Short: "",
	Long:  ``,
	Run:   viewLabelCmdAction,
}

func init() {
	rootCmd.AddCommand(viewLabelCmd)

	viewLabelCmd.Flags().IntP("board", "b", 0, "ID of the board")
	viewLabelCmd.Flags().IntSliceP("label", "l", []int{}, "ID of the label")

	viewLabelCmd.MarkFlagRequired("board")
	viewLabelCmd.MarkFlagRequired("label")
}
