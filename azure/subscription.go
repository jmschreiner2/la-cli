package azure

import (
	"context"
	"os"
	"strings"

	"github.com/jmschreiner2/la-cli/logger"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

const configSubscriptionKey = "subscriptionId"

type subscription struct {
	Name string
	ID   string
}

func promptSubscriptionID() string {
	logger := logger.NewLogger()
	logger.Debug("Loading azure credentials")
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		logger.Error("Could not load credentials.", logger.Args("Error", err))
		os.Exit(1)
	}
	factory, err := armsubscription.NewClientFactory(credential, nil)
	if err != nil {
		logger.Error("Failed to get subscription list.", logger.Args("Error", err))
		os.Exit(1)
	}

	logger.Debug("Getting list of subscriptions.")
	client := factory.NewSubscriptionsClient()
	subscriptions := client.NewListPager(nil)

	subList := make([]subscription, 0, 20)

	for subscriptions.More() {
		page, err := subscriptions.NextPage(context.TODO())
		if err != nil {
			logger.Error("Failed to load subscription ids.", logger.Args("Error", err))
			os.Exit(1)
		}

		for _, sub := range page.ListResult.Value {
			subList = append(subList, subscription{Name: strings.Clone(*sub.DisplayName), ID: strings.Clone(*sub.SubscriptionID)})
		}
	}

	if len(subList) == 0 {
		logger.Error("No Subscriptions To Select")
		os.Exit(0)
	}
	subID := ""
	if len(subList) == 1 {
		subID = subList[0].ID
		logger.Info("Auto selecting only subscription", logger.Args("Subscription", subID))
	} else {
		prompt := promptui.Select{
			Label: "Azure Subscription ID",
			Items: subList,
		}

		_, result, err := prompt.Run()
		if err != nil {
			logger.Error("Failed to select subscription id.", logger.Args("Error", err))
			os.Exit(1)
		}

		subID = result
	}

	return subID
}

func GetSubscriptionID() string {
	logger := logger.NewLogger()
	configVal := viper.GetString(configSubscriptionKey)

	if len(configVal) != 0 {
		logger.Debug("Using Subscription ID from Config", logger.Args("Subscription Id", configVal))
		return configVal
	}

	return promptSubscriptionID()
}

func SetSubscriptionID(id string) {
	logger := logger.NewLogger()
	if len(id) != 0 {
		logger.Debug("Updating Subscription ID to predetermined value.", logger.Args("subscription-id", id))
		viper.Set(configSubscriptionKey, id)
	}

	selectedID := promptSubscriptionID()
	logger.Debug("Updating Subscription ID to selected value.", logger.Args("subscription-id", selectedID))
	viper.Set(configSubscriptionKey, selectedID)
}
