package main

import (
	"fmt"

	"github.com/jonioliveira/go-template/pkg/app"
)

func main() {
	if err := app.Start(); err != nil {
		fmt.Printf("error while running app, %s", err)
	}

	fmt.Printf("Hello %s", "Base-Go")
}
