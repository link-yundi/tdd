package main

import (
	"github.com/link-yundi/tdd/ch08_dependency_injection"
	"os"
)

func main() {
	ch08_dependency_injection.Greet(os.Stdout, "ZhangYundi")
}
