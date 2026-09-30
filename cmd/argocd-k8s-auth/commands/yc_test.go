package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yandex-cloud/go-sdk/v2/credentials"
)

func TestYCCredentialsFromEnvironment(t *testing.T) {
	t.Run("uses metadata service by default", func(t *testing.T) {
		t.Setenv(envYCServiceAccountKeyFile, "")

		creds, err := ycCredentialsFromEnvironment()
		if err != nil {
			t.Fatalf("expected metadata credentials, got error: %v", err)
		}
		if _, ok := creds.(credentials.MetadataServiceCredentialProvider); !ok {
			t.Fatalf("expected metadata service credentials, got %T", creds)
		}
	})

	t.Run("loads authorized key file from environment", func(t *testing.T) {
		keyFile := filepath.Join(t.TempDir(), "missing-key.json")
		t.Setenv(envYCServiceAccountKeyFile, keyFile)

		_, err := ycCredentialsFromEnvironment()
		if err == nil {
			t.Fatal("expected an error for a missing authorized key file")
		}
		if _, statErr := os.Stat(keyFile); !os.IsNotExist(statErr) {
			t.Fatalf("test key file unexpectedly exists: %v", statErr)
		}
	})
}
