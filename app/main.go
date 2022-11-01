package main

import (
  "app/hello"
  "fmt"
)

func main() {
  t := hello.Hello()
  fmt.Printf("%v\n", t)
}
