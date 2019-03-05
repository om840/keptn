package credentialmanager

import (
	"os"
	"testing"

	"github.com/keptn/keptn/cli/utils"
)

const testEndPoint = "my-endpoint/"
const testAPIToken = "super-secret"

func init() {
	utils.Init(os.Stdout, os.Stdout, os.Stderr)
}

func TestSetAndGetCreds(t *testing.T) {

	if err := SetCreds(testEndPoint, testAPIToken); err != nil {
		t.Fatal(err)
	}

	endPoint, apiToken, err := GetCreds()
	if err != nil {
		t.Fatal(err)
	}
	if testEndPoint != endPoint || testAPIToken != apiToken {
		t.Fatal("Readed creds do not match")
	}
}