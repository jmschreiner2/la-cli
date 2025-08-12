package azure

import (
	"log/slog"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var credentials *azidentity.DefaultAzureCredential

func GetCredentials() *azidentity.DefaultAzureCredential {
	if credentials != nil {
		return credentials
	}

	slog.Debug("Loading azure credentials")
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		slog.Error("Could not load credentials.", "Error", err)
		os.Exit(1)
	}

	credentials = creds
	return creds
}
