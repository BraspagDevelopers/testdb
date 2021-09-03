package main

import (
	"fmt"
	"os"
	"strings"
)

func printResult(t Test, r TestResult, err error, namelen int) {
	var sb strings.Builder
	out := os.Stdout

	sb.WriteString(fmt.Sprintf("%s: ", padright(t.Name, namelen)))
	if err != nil {
		out = os.Stderr
		sb.WriteString("failed. ")
	} else {
		sb.WriteString("succeded")
	}
	sb.WriteString(fmt.Sprintf(" (%.3fs)", r.Duration.Seconds()))
	if err != nil {
		sb.WriteRune(' ')
		sb.WriteString(err.Error())
	}

	fmt.Fprintln(out, sb.String())
}
