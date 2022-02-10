package sendgrid

import (
	"os"
	"testing"
)

type TestContext struct {
	apiKey    string
	fromEmail string
}

func PrepareTest(t *testing.T) TestContext {
	apiKey, found := os.LookupEnv("SENDGRID_API_KEY")
	if !found {
		t.Fatal("Env var SENDGRID_API_KEY not specified")
	}

	fromEmail, found := os.LookupEnv("FROM_EMAIL")
	if !found {
		t.Fatal("Env var FROM_EMAIL not specified")
	}

	return TestContext{
		apiKey:    apiKey,
		fromEmail: fromEmail,
	}
}
