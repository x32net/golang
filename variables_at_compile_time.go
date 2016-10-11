// https://medium.com/@joshroppo/setting-go-1-5-variables-at-compile-time-for-versioning-5b30a965d33e#.i8utrlxxx
package main

import “fmt”
var MainVar string

func main() {
 fmt.Printf(“MainVar: %s\n”, MainVar)
}

// go build -ldflags "-X main.MainVar=hihi"
// ./bin
// MainVar: hihi
