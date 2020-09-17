package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type encoding int

func main() {
	var lines []string

	var decode bool
	flag.BoolVar(&decode, "d", false, "url decode")

	var times int
	flag.IntVar(&times, "t", 1, "encode/decode for x times")

	var recursive bool
	flag.BoolVar(&recursive, "r", false, "recursively decode until no more changes")

	flag.Parse()

	if times < 1 {
		times = 1
	}

	if flag.NArg() > 0 {
		lines = []string{flag.Arg(0)}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
	}

	for _, line := range lines {
		if decode {
			decoded := unescape(line)
			for i := 1; i < times; i++ {
				decoded = unescape(decoded)
			}
			if recursive {
				for decoded != unescape(decoded) {
					decoded = unescape(decoded)
				}
			}
			fmt.Println(decoded)
		} else {
			encoded := url.PathEscape(line)
			for i := 1; i < times; i++ {
				encoded = url.PathEscape(encoded)
			}
			fmt.Println(encoded)
		}
	}
}

// Copied and rewrote from https://golang.org/src/net/url/url.go since I want it to ignore invalid % encoding
func unescape(s string) string {
	n := 0

	var t strings.Builder
	t.Grow(len(s) - 2*n)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '%':
			if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
				t.WriteByte(s[i])
			} else {
				t.WriteByte(unhex(s[i+1])<<4 | unhex(s[i+2]))
				i += 2
			}
		default:
			t.WriteByte(s[i])
		}
	}
	return t.String()
}

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
