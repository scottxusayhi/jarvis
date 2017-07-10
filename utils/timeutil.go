package utils

import (
	"bytes"
	"bufio"
	"fmt"
	"time"
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

func ISO8601Now() (time.Time, string) {
	now := time.Now()
	return now, now.Format("2006-01-02T15:04:05Z0700")
}
