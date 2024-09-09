package main

import (
	"fmt"

	"github.com/imirjar/rb-auth/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
