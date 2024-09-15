/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"godeck/deck"

	"github.com/spf13/cobra"
)

func deleteACLCmdAction(cmd *cobra.Command, args []string) {
	boardId, _ := cmd.Flags().GetInt("board")
	aclId, _ := cmd.Flags().GetInt("acl")


  r, err := deck.DeleteACLRule(boardId, aclId)
  if err != nil {
    fmt.Println("There was an error: ", err)
    return
  }

  fmt.Println(r)
}

// deleteACLCmd represents the deleteACL command
var deleteACLCmd = &cobra.Command{
	Use:   "deleteACL",
	Short: "",
	Long:  ``,
	Run:   deleteACLCmdAction,
}

func init() {
	rootCmd.AddCommand(deleteACLCmd)

	deleteACLCmd.Flags().IntP("board", "b", -1, "Board ID")
	deleteACLCmd.Flags().IntP("acl", "a", -1, "ACL ID")

	deleteACLCmd.MarkFlagRequired("board")
	deleteACLCmd.MarkFlagRequired("acl")
}

