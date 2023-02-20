package main

import "strings"

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 se "/" não for encontrada
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
