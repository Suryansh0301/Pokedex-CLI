package main

import (
	"slices"
	"strings"
)

func cleanInput(text string) []string {
	return slices.DeleteFunc(strings.Split(strings.TrimSpace(text), " "), func(item string) bool {
		return item == ""
	})

}
