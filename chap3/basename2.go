package main

func basename(s string) string {
	slash := string.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := string.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}