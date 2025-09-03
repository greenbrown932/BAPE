// cmd/api/main.go
package main

import (
	"BAPE/internal/api"
)

func main() {
	r := api.SetupRouter()
	r.Run(":8080")
}
