/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jmschreiner2/la-cli/azure"
	"github.com/spf13/cobra"
)

// subscriptionCmd represents the subscription command
var subscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Args:  cobra.ArbitraryArgs,
	Short: "Configure the Azure subscription used by la-cli",
	Long: `Configure the Azure Subscription ID that la-cli will use when querying Logic App information.

You can optionally provide a subscription ID as an argument. After setting it, la-cli will open an interactive
selector so you can confirm or choose a different subscription from your account. If you do not provide an ID,
la-cli will directly open the interactive selector. If your account has only one subscription, it will be
selected automatically.

Usage:
  la-cli set subscription [subscription-id]
`,
	Run: func(cmd *cobra.Command, args []string) {
		subID := ""
		if len(args) > 0 {
			subID = args[0]
		}

		azure.SetSubscriptionID(subID)
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
