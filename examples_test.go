package template_test

import (
	"atomicgo.dev/template"
	"fmt"
)

func Example_demo() {
	fmt.Println(template.HelloWorld())
	// Output: Hello, World!
}

func ExampleHelloWorld() {
	fmt.Println(template.HelloWorld())
	// Output: Hello, World!
}
