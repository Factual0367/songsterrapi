# Songster API

This is a Go module that enables searching for tabs on Songsterr and get download links for Guitar Pro files.

## Installation

```bash
go get github.com/onurhanak/songsterrapi
```

## Usage

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/onurhanak/songsterrapi"
)

func main() {
	router := gin.Default()
	router.GET("/search", songsterrapi.SearchRequest)
	router.Run("localhost:8080")
}
```

Then you can use any tool you want to make a request and get results:

```sh
http://127.0.0.1:8080/search?query={ExampleQuery}
```
