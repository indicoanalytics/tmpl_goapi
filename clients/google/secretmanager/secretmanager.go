package secretmanager

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/entity"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/iterator"
)

type GCPSecretManager struct{}

func New() *GCPSecretManager {
	return &GCPSecretManager{}
}

func newClient() (context.Context, *secretmanager.Client) {
	ctx := context.Background()

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error while trying to connect to google cloud secret manager",
			Reason:  err.Error(),
		}, "critical", nil)

		panic(err)
	}

	return ctx, client
}

func (secretmanager *GCPSecretManager) ListSecrets(parent string, filterPrefix string) map[string]interface{} {
	ctx, client := newClient()
	defer client.Close()

	filter := filterPrefix
	if filter == "" {
		filter = "*"
	}

	data := &secretmanagerpb.ListSecretsRequest{
		Parent: fmt.Sprintf("projects/%s", parent),
		Filter: fmt.Sprintf("Name: %s", filter),
	}

	secretList := map[string]interface{}{}
	secrets := client.ListSecrets(ctx, data)

	for {
		secret, err := secrets.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			logging.Log(&entity.LogDetails{
				Message: "error to get next secret in google secret manager",
				Reason:  err.Error(),
			}, "critical", nil)

			panic(err)
		}

		splitSecret := strings.Split(secret.Name, "/")
		secretName := splitSecret[len(splitSecret)-1]

		if filterPrefix != "" {
			if strings.Contains(secretName, filterPrefix) {
				secretName = strings.Split(secretName, filterPrefix)[1]
			}
		}

		secretList[secretName] = secretmanager.accessSecretVersion(fmt.Sprintf("%s/versions/latest", secret.Name))
	}

	return secretList
}

func (secretmanager *GCPSecretManager) accessSecretVersion(version string) string {
	ctx, client := newClient()
	defer client.Close()

	result, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: version,
	})
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error to access secret version in google secret manager",
			Reason:  err.Error(),
		}, "critical", nil)

		panic(err)
	}

	return string(result.Payload.Data)
}
