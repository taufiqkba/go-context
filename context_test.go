package gocontext

import (
	"context"
	"fmt"
	"testing"
)

// Context Background
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
