/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func updateACLCmdAction(cmd *cobra.Command, args []string) {
	nType, _ := cmd.Flags().GetInt("type")
  nParticipant, _ := cmd.Flags().GetString("participant")
  permEdit, _ := cmd.Flags().GetBool("permedit")
  permShare, _ := cmd.Flags().GetBool("permShare")
  permManage, _ := cmd.Flags().GetBool("permManage")

	boardId, _ := cmd.Flags().GetInt("board")
	aclId, _ := cmd.Flags().GetInt("acl")

	payload := map[string]interface{}{
		"type":  nType,
    "participant": nParticipant,
    "permissionEdit": permEdit,
    "permissionShare": permShare,
    "permissionManage": permManage,
	}

  r, err := deck.UpdateACLRule(boardId, aclId, payload)
  if err != nil {
    fmt.Println("There was an error: ", err)
    return
  }

  fmt.Println(r)
}

// updateACLCmd represents the updateACL command
var updateACLCmd = &cobra.Command{
	Use:   "updateACL",
	Short: "",
	Long:  ``,
	Run:   updateACLCmdAction,
}

func init() {
	rootCmd.AddCommand(updateACLCmd)

	updateACLCmd.Flags().IntP("board", "b", -1, "Board ID")
	updateACLCmd.Flags().IntP("acl", "a", -1, "ACL ID")

	updateACLCmd.Flags().IntP("type", "t", 0, "Type of participant")
	updateACLCmd.Flags().StringP("participant", "p", "", "The UID of the participant")
	updateACLCmd.Flags().BoolP("permedit", "e", false, "Permission to edit")
	updateACLCmd.Flags().BoolP("permshare", "s", false, "Permission to share")
	updateACLCmd.Flags().BoolP("permmanaged", "m", false, "Permission to manage")

	updateACLCmd.MarkFlagRequired("board")
	updateACLCmd.MarkFlagRequired("acl")
	updateACLCmd.MarkFlagRequired("type")
	updateACLCmd.MarkFlagRequired("participant")
}


