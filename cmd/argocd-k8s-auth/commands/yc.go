package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"

	"github.com/argoproj/argo-cd/v3/util/errors"
)

const envYCServiceAccountKeyFile = "YC_SERVICE_ACCOUNT_KEY_FILE"

func newYCCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "yc",
		Short: "Generate a Yandex Cloud IAM token",
		Run: func(c *cobra.Command, _ []string) {
			creds, err := ycCredentialsFromEnvironment()
			errors.CheckError(err)
			sdk, err := ycsdk.Build(c.Context(), options.WithCredentials(creds))
			errors.CheckError(err)
			token, err := sdk.CreateIAMToken(c.Context())
			errors.CheckError(err)
			_, _ = fmt.Fprint(os.Stdout, formatJSON(token.GetIamToken(), token.GetExpiresAt()))
		},
	}
	return command
}

func ycCredentialsFromEnvironment() (credentials.Credentials, error) {
	if keyFile := os.Getenv(envYCServiceAccountKeyFile); keyFile != "" {
		return credentials.ServiceAccountKeyFile(keyFile)
	}
	return credentials.MetadataService(), nil
}
