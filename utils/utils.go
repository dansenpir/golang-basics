package utils

import (
	"fmt"
	"strings"
)

func PrintSection(title string) {
	fmt.Printf("\n%s\n%s\n", title, strings.Repeat("-", len(title)))
}
