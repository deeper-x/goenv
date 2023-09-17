### Golang Dot Env

Handle .env parameters in golang.

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

Accessing non existing variables
```golang
efo, err := New()
if err != nil {
    panic(err)
}

v, err = efo.Get("NON_EXISTING")
if err == nil {
    panic("handle error on non-existing key fetched")
}
```

Install:
```sh
go get github.com/deeper-x/golang-dot-env
```
