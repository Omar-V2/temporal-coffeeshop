package wf

import (
	"context"
	"fmt"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}
