# go-memory-cache

For install

``go get github.com/bald-cat/go-memory-cache``

Usage example 

```go
package main

import (
	"fmt"
	cache2 "github.com/bald-cat/go-memory-cache"
)

func main() {
	cache := cache2.NewCache()
	cache.Set("userId", 42, 60)
	fmt.Println(cache.Get("userId"))

	author := struct {
		name  string
		login string
		age   int
	}{
		name:  "Ihor",
		login: "bald-cat",
		age:   15,
	}

	cache.Set("author", author, 60)
	fmt.Println(cache.Get("author"))
}
```