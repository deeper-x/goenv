### Golang Dot Env

Handle .env parameters in golang.

Usage:

Step 1 - Install:
```sh
go get -v github.com/deeper-x/golang-dot-env@latest
```

Step 2 - Create an .env:
```txt
USER=demo_user
PASSWORD=qwerty
```

Step 3 - Deploy:
```golang
package main

import (
	"fmt"

	gde "github.com/deeper-x/golang-dot-env"
)


func main() {
	cfo, err := gde.New()
	if err != nil {
		panic(err)
	}

	username, err := cfo.Get("USERNAME")
	if err != nil {
		panic(err)
	}

	password, err := cfo.Get("PASSWORD")
	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %s; password: %s", username, password)
}

// Output:
// user: demo_user; password: qwerty
```

Test:
```golang
go test -v ./...
```

### Explanation:

Step 1: build file object
```golang
efo, err := New() // <-- access default .env file in current dir
// efo, err := New("/path/to/other/file/.env") <-- Or pass a custom path
if err != nil {
	panic(err)
}
```

Step 2: access .env parameters
```golang
v, err = efo.Get("USERNAME")
if err == nil {
    panic("...handle error...")
}

fmt.Println(v) // <-- access environment value
```

PLEASE NOTE: Accessing non existing variables raises an error!
```golang
efo, err := New()
if err != nil {
    panic(err)
}

v, err = efo.Get("NON_EXISTING")
if err != nil {
    panic("handle error on non-existing key fetched")
}
```

