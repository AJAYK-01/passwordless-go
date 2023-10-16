# Introduction
Community built Passwordless.dev sdk library for Go  🚀

Official Passwordless docs: https://docs.passwordless.dev

Admin portal to create an account: https://admin.passwordless.dev

[![Go Reference](https://pkg.go.dev/badge/github.com/AJAYK-01/passwordless-go/passwordless.svg)](https://pkg.go.dev/github.com/AJAYK-01/passwordless-go/passwordless)

# Installation

``` 
go get https://github.com/AJAYK-01/passwordless-go
```

# Usage

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

