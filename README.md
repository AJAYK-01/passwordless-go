### Winner at [Passwordless.dev Hackathon](https://community.bitwarden.com/t/passwordless-dev-hackathon-winners-announcement/61281) !

# Introduction
Community built Passwordless.dev sdk library for Go  🚀

Official Passwordless docs: https://docs.passwordless.dev

Admin portal to create an account: https://admin.passwordless.dev

[![Go Reference](https://pkg.go.dev/badge/github.com/AJAYK-01/passwordless-go/passwordless.svg)](https://pkg.go.dev/github.com/AJAYK-01/passwordless-go/passwordless)


# Installation

``` 
go get https://github.com/AJAYK-01/passwordless-go
```

# Import

Importing package and creating client
```Go
import (
	"github.com/AJAYK-01/passwordless-go/passwordless"
)

func main() {
    pClient := passwordless.Client{ApiSecret: "API_SECRET", BaseUrl: "v4.passwordless.dev"}
}
```

Refer the functions Docs [Here](https://pkg.go.dev/github.com/AJAYK-01/passwordless-go/passwordless#pkg-functions)

# Example

```go
package main

import (
	"fmt"

	"github.com/AJAYK-01/passwordless-go/passwordless"
)

func main() {
	api := passwordless.Client{ApiSecret: "myapplication:secret:11f8dd7733744f2596f2a28544b5fbc4", BaseUrl: "v4.passwordless.dev"}

	// Create a registration token
	params := passwordless.RegisterRequest{UserId: "test", Username: "test", Displayname: "Test User"}
	resp, err := api.CreateRegisterToken(params)
	if err != nil {
		fmt.Printf("got an error: %v", err)
	}
	fmt.Printf("Token: %s\n", resp.Token)

	// Verify a sign-in
	resp2, err := api.VerifySignin("verify_d5vzCkL_GvpS4VYtoT3hbsdkghbskdgbksjbdgksjbgkb")
	if err != nil {
		fmt.Printf("got an error: %v", err)
	}
	fmt.Printf("Success: %t\n", resp2.Success)

	// Set aliases
	existingAliases := []string{"alias1", "alias2"}
	aliases := append(existingAliases, "helloalias1", "helloalias2")

	params2 := passwordless.AliasRequest{UserId: "40217fd0-52c6-4656-b617-f158fe8e9ed7", Aliases: aliases}
	err = api.SetAliases(params2)
	if err != nil {
		fmt.Printf("got an error: %v", err)
	}

	// List credentials
	resp3, err := api.ListCredentials("40217fd0-52c6-4656-b617-f158fe8e9ed7")
	if err != nil {
		fmt.Printf("got an error: %v", err)
	}
	fmt.Printf("Credentials: %v\n", resp3.Values)

	// Delete a credential
	params3 := passwordless.DeleteCredentialRequest{CredentialId: "8WSezHAkYT/Wez6drClajZ+sS5ruZLt93iSTcQJuPN0="}
	err = api.DeleteCredential(params3)
	if err != nil {
		fmt.Printf("got an error: %v", err)
	}
}
```

# Testing
Set env required to run tests, modify the env values to real/fake ones to check working
| Variable | Example Value |
| --- | --- |
| `API_SECRET` | `"myapplication:secret:11f8dd7733744f2596f2a28544b5fbc4"` |
| `VERIFY_TOKEN` | `"verify_d5vzCkL_GvpS4VYtoT3hbsdkghbskdgbksjbdgksjbgkb"` |
| `USER_ID` | `"40217fd0-52c6-4656-b617-f158fe8e9ed7"` |
| `EXISTING_ALIASES` | `"hello@hello.com"` |
| `LIST_CREDENTIAL_ID` | `"rpaeZKWIkyDi2sg4EQA0Jwlm0RTaWbAim8pd3pw8eec="` |
| `DELETE_CREDENTIAL_ID` | `"8WSezHAkYT/Wez6drClajZ+sS5ruZLt93iSTcQJuPN0="` |


Change to test directory
```
cd test
```
Run tests
```
go test
```

