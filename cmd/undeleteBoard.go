/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	// "strconv"
	//
 //  "godeck/deck"

	"github.com/spf13/cobra"
)

func undeleteBoardCmdAction(cmd *cobra.Command, args []string) {
	// ids := make([]int, len(args)-1)
	// for _, v := range args {
	// 	id, err := strconv.Atoi(v)
	// 	if err != nil {
	// 		fmt.Printf("Could not read id [%s]: %s\n", v, err)
	// 		continue
	// 	}
	// 	ids = append(ids, id)
	// }
	//
	// failed := deck.UndoDeleteBoards(ids)
	// if len(failed) == 0 {
	// 	fmt.Println("Boards recovered successfully.")
	// 	return
	// }
	//
	// // Show the user the errors that did happen
	// for k, v := range failed {
	// 	fmt.Printf("%-3d - %s\n", k, v)
	// }
}

// undeleteBoardCmd represents the undeleteBoard command
var undeleteBoardCmd = &cobra.Command{
	Use:   "undeleteBoard",
	Short: "",
	Long:  ``,
	Run:   undeleteBoardCmdAction,
}

func init() {
	// rootCmd.AddCommand(undeleteBoardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undeleteBoardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undeleteBoardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
