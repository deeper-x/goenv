### Golang Dot Env

Handle .env parameters in golang.

Usage:

Step 1 - Install:
```sh
go get -v github.com/deeper-x/goenv@latest
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
	gde "github.com/deeper-x/goenv"
)


func main() {
	cfo, err := gde.New(".env")
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

### Details

Build file object:
```golang
efo, err := gde.New() // <-- access default .env file in current dir, or New("./assets/.env") for custom file path
if err != nil {
	panic(err)
}
```
Inspect file content:
```golang
fc, err := FileDump()
if err != nil {
	panic(err)
}

fmt.Println(fc)
```

Access .env parameters:
```golang
v, err = cfo.Get("USERNAME")
if err == nil {
    panic("...handle error...")
}

fmt.Println(v) // <-- access environment value
```



PLEASE NOTE: Accessing non existing variables raises an error!
```golang
efo, err := goenv.New()
if err != nil {
    panic(err)
}

v, err = efo.Get("NON_EXISTING")
if err != nil {
    panic("handle error on non-existing key fetched")
}
```

