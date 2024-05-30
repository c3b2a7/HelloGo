package main

import (
	"context"
	"github.com/1password/onepassword-sdk-go"
	"os"
)

// see https://github.com/1Password/onepassword-sdk-go/tree/main?tab=readme-ov-file#-get-started
func main() {
	token := os.Getenv("OP_SERVICE_ACCOUNT_TOKEN")

	client, err := onepassword.NewClient(
		context.TODO(),
		onepassword.WithServiceAccountToken(token),
		onepassword.WithIntegrationInfo("My 1Password Integration", "v1.0.0"),
	)
	if err != nil {
		// handle err
	}
	secret, err := client.Secrets.Resolve(context.Background(), "op://vault/item/field")
	if err != nil {
		// handle err
	}
	// do something with the secret
	println(secret)
}
