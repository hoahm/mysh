package util

import (
	"fmt"
)

// ServerKey returns the key as string
func ServerKey(p string, n int) string {
	return fmt.Sprintf("projects.%s.%d", p, n)
}
