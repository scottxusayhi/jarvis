package utils

import (
	"bytes"
	"bufio"
	"fmt"
)

func FormatSecond(second uint64) string {
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)

	days := second / (60 * 60 * 24)

	if days != 0 {
		s := ""
		if days > 1 {
			s = "s"
		}
		fmt.Fprintf(w, "%d day%s, ", days, s)
	}

	minutes := second / 60
	hours := minutes / 60
	hours %= 24
	minutes %= 60

	fmt.Fprintf(w, "%2d:%02d", hours, minutes)

	w.Flush()
	return buf.String()
}
