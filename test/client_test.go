package passwordless

import (
	"os"
	"strings"
	"testing"

	"github.com/AJAYK-01/passwordless-go/passwordless"
)

var api = passwordless.Client{ApiSecret: os.Getenv("API_SECRET"), BaseUrl: "v4.passwordless.dev"}

func TestCreateRegisterToken(t *testing.T) {
	params := passwordless.RegisterRequest{UserId: "test", Username: "test", Displayname: "Test User"}
	resp, err := api.CreateRegisterToken(params)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	if !strings.HasPrefix(resp.Token, "register_") {
		t.Errorf("Token does not start with 'register_': %s", resp.Token)
	}
}

func TestVerifySignin(t *testing.T) {
	resp, err := api.VerifySignin(os.Getenv("VERIFY_TOKEN"))
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	if resp.Success != true {
		t.Errorf("got: %t, want: true", resp.Success)
	}
}

func TestSetAliases(t *testing.T) {
	existingAliases := strings.Split(os.Getenv("EXISTING_ALIASES"), ",")
	aliases := append(existingAliases, "helloalias1", "helloalias2")

	params := passwordless.AliasRequest{UserId: os.Getenv("USER_ID"), Aliases: aliases}
	err := api.SetAliases(params)

	if err != nil {
		t.Errorf("got an error: %v", err)
	}
}

func TestListCredentials(t *testing.T) {
	resp, err := api.ListCredentials(os.Getenv("USER_ID"))
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	if len(resp.Values) == 0 || resp.Values[0].Descriptor.Type != "public-key" || resp.Values[0].Descriptor.Id != os.Getenv("LIST_CREDENTIAL_ID") {
		t.Errorf("got unexpected response: %v", resp.Values)
	}
}

func TestDeleteCredential(t *testing.T) {
	params := passwordless.DeleteCredentialRequest{CredentialId: os.Getenv("DELETE_CREDENTIAL_ID")}
	err := api.DeleteCredential(params)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
}
