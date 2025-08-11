package azure

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

func SelectLogicApp() {
	slog.Debug("Loading azure credentials")
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		slog.Error("Could not load credentials.", "Error", err)
		os.Exit(1)
	}
	subID := GetSubscriptionID()
	factory, err := armlogic.NewClientFactory(subID, credential, nil)
	if err != nil {
		slog.Error("Failed to get logic app list.", "Error", err)
		os.Exit(1)
	}

	slog.Debug("Getting list of logic app workflows.")
	client := factory.NewWorkflowsClient()
	workflows := client.NewListBySubscriptionPager(nil)
	workflowList := make([]*armlogic.Workflow, 0, 20)

	for workflows.More() {
		page, err := workflows.NextPage(context.TODO())
		if err != nil {
			slog.Error("Failed to load logic apps.", "Error", err)
			os.Exit(1)
		}

		for _, workflow := range page.WorkflowListResult.Value {
			fmt.Println(*workflow.Name, *workflow.Location, *workflow.Type)
			workflowList = append(workflowList, workflow)
		}
	}

	if len(workflowList) == 0 {
		slog.Error("No Logic Apps To Select")
		os.Exit(0)
	}
	/*_:= promptui.Select{
		Label: "Azure Workflows",
		Items: workflowList,
	}

	_, result, err := prompt.Run()
	if err != nil {
		slog.Error("Failed to select subscription id.", "Error", err)
		os.Exit(1)
	}*/

	// return subID
}
