package hello

import "fmt"

func Hello() string {
  return "Hello World!!"
}

func HelloWithName(name string) string {
  return fmt.Sprintf("Hello, %v!!", name)
}
