package main

import (
	"fmt"
	"strconv"
)

func coalesce(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

func padright(value string, length int) string {
	template := "%-" + strconv.Itoa(length) + "s"
	return fmt.Sprintf(template, value)
}
