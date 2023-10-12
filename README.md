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

Example usage of methods in [client_test.go](test/client_test.go)
