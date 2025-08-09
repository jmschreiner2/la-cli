/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subscriptionCmd represents the subscription command
var subscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Args:  cobra.ArbitraryArgs,
	Short: "Set your Azure Subscription ID",
	Long: `Set you Azure Subscription ID
usage: la-cli set subscription [subsciption-id]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subscription called")
	},
}

func init() {
	setCmd.AddCommand(subscriptionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscriptionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscriptionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
