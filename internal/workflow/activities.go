// internal/workflow/activities.go
package workflow

import (
	"context"
	"fmt"
)

func BapeActivity(ctx context.Context, name string) (string, error) {
	fmt.Printf("Running activity with input: %s\n", name)
	return "Hello, " + name + "!", nil
}
