package main

import (
	"context"
	"fmt"

	"github.com/imirjar/rb-auth/internal/app"
)

func main() {
	ctx := context.Background()
	if err := app.Run(ctx); err != nil {
		fmt.Println(err)
	}
}
