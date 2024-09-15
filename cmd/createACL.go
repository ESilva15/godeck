/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func createACLCmdAction(cmd *cobra.Command, args []string) {
	nType, _ := cmd.Flags().GetInt("type")
  nParticipant, _ := cmd.Flags().GetString("participant")
  permEdit, _ := cmd.Flags().GetBool("permedit")
  permShare, _ := cmd.Flags().GetBool("permShare")
  permManage, _ := cmd.Flags().GetBool("permManage")

	boardId, _ := cmd.Flags().GetInt("board")

	payload := map[string]interface{}{
		"type":  nType,
    "participant": nParticipant,
    "permissionEdit": permEdit,
    "permissionShare": permShare,
    "permissionManage": permManage,
	}

  r, err := deck.CreateACLRule(boardId, payload)
  if err != nil {
    fmt.Println("There was an error: ", err)
    return
  }

  fmt.Println(r)
}

// createACLCmd represents the createACL command
var createACLCmd = &cobra.Command{
	Use:   "createACL",
	Short: "",
	Long:  ``,
	Run:   createACLCmdAction,
}

func init() {
	rootCmd.AddCommand(createACLCmd)

	createACLCmd.Flags().IntP("board", "b", -1, "Board in which to insert the new card.")

	createACLCmd.Flags().IntP("type", "t", 0, "Type of participant")
	createACLCmd.Flags().StringP("participant", "p", "", "The UID of the participant")
	createACLCmd.Flags().BoolP("permedit", "e", false, "Permission to edit")
	createACLCmd.Flags().BoolP("permshare", "s", false, "Permission to share")
	createACLCmd.Flags().BoolP("permmanaged", "m", false, "Permission to manage")

	createACLCmd.MarkFlagRequired("board")
	createACLCmd.MarkFlagRequired("type")
	createACLCmd.MarkFlagRequired("participant")
}

