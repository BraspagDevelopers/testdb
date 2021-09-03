package main

import (
	"fmt"
	"strconv"
)

func coalesce(value string, fallback ...string) string {
	if value == "" && len(fallback) > 0 {
		coalesce(fallback[0], fallback[1:]...)
	}
	return value
}

func padright(value string, length int) string {
	template := "%-" + strconv.Itoa(length) + "s"
	return fmt.Sprintf(template, value)
}
