package azure

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/jmschreiner2/la-cli/logger"
)

var credentials *azidentity.DefaultAzureCredential

func GetCredentials() *azidentity.DefaultAzureCredential {
	logger := logger.NewLogger()
	if credentials != nil {
		return credentials
	}

	logger.Debug("Loading azure credentials")
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		logger.Error("Could not load credentials.", logger.Args("Error", err))
		os.Exit(1)
	}

	credentials = creds
	return creds
}
