package azure

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
	"github.com/manifoldco/promptui"
)

const logicAppKind = "functionapp,workflowapp"

func getLogicAppList() []*armappservice.Site {
	slog.Info("Getting list of Logic Apps")
	slog.Debug("Initializing Factory")
	creds := GetCredentials()
	subID := GetSubscriptionID()

	factory, err := armappservice.NewClientFactory(subID, creds, nil)
	if err != nil {
		slog.Error("Failed to get logic app list.", "Error", err)
		os.Exit(1)
	}

	slog.Debug("Getting logic app pager")
	client := factory.NewWebAppsClient()
	logicapps := client.NewListPager(nil)

	selectList := make([]*armappservice.Site, 0, 5)
	for logicapps.More() {
		page, err := logicapps.NextPage(context.TODO())
		if err != nil {
			slog.Error("Failed to load logic apps.", "Error", err)
			os.Exit(1)
		}

		for _, app := range page.WebAppCollection.Value {
			if logicAppKind == *app.Kind {
				selectList = append(selectList, app)
			}
		}
	}

	return selectList
}

type logicAppItem struct {
	Name          string
	ResourceGroup string
	LogicApp      *armappservice.Site
}

func promptLogicAppSelector(logicApps []*armappservice.Site) *armappservice.Site {
	slog.Debug("Setting up select for logic apps", "Logic App Count", len(logicApps))
	if len(logicApps) == 0 {
		return nil
	} else if len(logicApps) == 1 {
		return logicApps[0]
	}

	items := make([]logicAppItem, len(logicApps))

	for i, app := range logicApps {
		items[i] = logicAppItem{
			Name:          strings.Clone(*app.Name),
			ResourceGroup: strings.Clone(*app.Properties.ResourceGroup),
			LogicApp:      app,
		}
	}

	template := &promptui.SelectTemplates{
		Label: "{{ . }}?",
		// Active:   fmt.Sprintf(`%s {{ .Name | underline }}{{ " (" | underline }}{{ .ResourceGroup | underline }}{{ ")" | underline }}`, promptui.IconSelect),
		Active:   fmt.Sprintf(`%s {{ printf "%%s (%%s)" .Name .ResourceGroup | underline }}`, promptui.IconSelect),
		Inactive: " {{ .Name }} ({{ .ResourceGroup }})",
		Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .Name | faint }} ({{ .ResourceGroup | faint }})`, promptui.IconGood),
	}

	prompt := promptui.Select{
		Label:     "Select Logic App",
		Items:     items,
		Templates: template,
	}

	i, _, err := prompt.Run()
	if err != nil {
		slog.Error("Failed to reder logic app select.", "Error", err)
		os.Exit(1)
	}

	return items[i].LogicApp
}

func SelectLogicApp() *armappservice.Site {
	slog.Debug("Loading azure credentials")
	credentials := GetCredentials()
	subID := GetSubscriptionID()

	logicApps := getLogicAppList()

	return promptLogicAppSelector(logicApps)
}

func getWorkflowList(creds *azidentity.DefaultAzureCredential, subID string, logicApp *armappservice.Site) []*armappservice.WorkflowEnvelope {
	slog.Info("Getting list of Workflows")
	slog.Debug("Initializing Factory")
	factory, err := armappservice.NewClientFactory(subID, creds, nil)
	if err != nil {
		slog.Error("Failed to get workflow list.", "Error", err)
		os.Exit(1)
	}

	client := factory.NewWebAppsClient()
	workflows := client.NewListWorkflowsPager(*logicApp.Properties.ResourceGroup, *logicApp.Name, nil)

	selectList := make([]*armappservice.WorkflowEnvelope, 0, 10)
	for workflows.More() {
		page, err := workflows.NextPage(context.TODO())
		if err != nil {
			slog.Error("Failed to load logic apps.", "Error", err)
			os.Exit(1)
		}

		for _, flow := range page.WorkflowEnvelopeCollection.Value {
			fmt.Println(*flow.Name, *flow.Properties.FlowState, *flow.ID)
			selectList = append(selectList, flow)
		}
	}

	return selectList
}

func getWorkflowDetails(resourceGroup string, siteName string, workflowName string) {
	slog.Info("Getting workflow details")
	slog.Debug("Initializing Factory")
	creds := GetCredentials()
	subID := GetSubscriptionID()
	factory, err := armappservice.NewClientFactory(subID, creds, nil)
	if err != nil {
		slog.Error("Failed to get workflow list.", "Error", err)
		os.Exit(1)
	}

	client := factory.NewWebAppsClient()
	workflow, err := client.GetWorkflow(context.TODO(), resourceGroup, siteName, workflowName, nil)

	return workflow
}

func SelectWorkflow(logicApp *armappservice.Site) *armappservice.Workflow {
	slog.Debug("Loading azure credentials")
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		slog.Error("Could not load credentials.", "Error", err)
		os.Exit(1)
	}
	subID := GetSubscriptionID()

	workflows := getWorkflowList(credential, subID, logicApp)

	return workflows[0]
}
