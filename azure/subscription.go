package subscription

import (
	"log/slog"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func promptSubscriptionId() {
	prompt := promptui.Select {
		Label: "Azure Subscription ID",
		Items: [],
	}
}

func GetSubscriptionId() string {
	configVal := viper.GetString("subscriptionId")

	if len(configVal) != 0 {
		slog.Debug("Using Subscription ID from Config", "subscription-id", configVal)
		return configVal
	}

	prompt := promptui.Prompt{
		Label: "Azure Subscription ID",
	}
	result, err := prompt.Run()
	if err != nil {
		return ""
	}
	return result
}

func SetSubscriptionId(id string) {
}
