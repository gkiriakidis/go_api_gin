# go_api_gin

This is an simple CRUD example to show how to use Gin as a router.
So far, without the use of a database.
For more detailed CRUD example (with Gorilla/Mux router) check: https://github.com/mnikolud/go_api

## Installation of Gin Web Framework

To install Gin package, you need to install Go and set your Go workspace first.

1. Download and install it:

```sh
$ go get -u github.com/gin-gonic/gin
```

2. Import it in your code:

```go
import "github.com/gin-gonic/gin"
```

3. (Optional) Import `net/http`. This is required for example if using constants such as `http.StatusOK`.

```go
import "net/http"
```

## Other?
