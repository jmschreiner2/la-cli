package azure

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

const configSubscriptionKey = "subscriptionId"

type subscription struct {
	Name string
	Id   string
}

func promptSubscriptionID() string {
	slog.Debug("Loading azure credentials")
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		slog.Error("Could not load credentials.", "Error", err)
		os.Exit(1)
	}
	factory, err := armsubscription.NewClientFactory(credential, nil)
	if err != nil {
		slog.Error("Failed to get subscription list.", "Error", err)
		os.Exit(1)
	}

	slog.Debug("Getting list of subscriptions.")
	client := factory.NewSubscriptionsClient()
	subscriptions := client.NewListPager(nil)

	subList := make([]subscription, 0, 20)

	for subscriptions.More() {
		page, err := subscriptions.NextPage(context.TODO())
		if err != nil {
			slog.Error("Failed to load subscription ids.", "Error", err)
			os.Exit(1)
		}

		for _, sub := range page.ListResult.Value {
			subList = append(subList, subscription{Name: strings.Clone(*sub.DisplayName), Id: strings.Clone(*sub.SubscriptionID)})
		}
	}

	if len(subList) == 0 {
		slog.Error("No Subscriptions To Select")
		os.Exit(0)
	}
	subID := ""
	if len(subList) == 1 {
		subID = subList[0].Id
		slog.Info("Auto selecting only subscription", "Subscription", subID)
	} else {
		prompt := promptui.Select{
			Label: "Azure Subscription ID",
			Items: subList,
		}

		_, result, err := prompt.Run()
		if err != nil {
			slog.Error("Failed to select subscription id.", "Error", err)
			os.Exit(1)
		}

		subID = result
	}

	return subID
}

func GetSubscriptionID() string {
	configVal := viper.GetString(configSubscriptionKey)

	if len(configVal) != 0 {
		slog.Debug("Using Subscription ID from Config", "subscription-id", configVal)
		return configVal
	}

	return promptSubscriptionID()
}

func SetSubscriptionID(id string) {
	if len(id) != 0 {
		slog.Debug("Updating Subscription ID to predetermined value.", "subscription-id", id)
		viper.Set(configSubscriptionKey, id)
	}

	selectedID := promptSubscriptionID()
	slog.Debug("Updating Subscription ID to selected value.", "subscription-id", selectedID)
	viper.Set(configSubscriptionKey, selectedID)
}
