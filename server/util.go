package main

import "os"
import "utf8"

func isText(b []byte) bool {
	for len(b) > 0 && utf8.FullRune(b) {
		if rune, size := utf8.DecodeRune(b); size == 1 && rune == utf8.RuneError {
			return false
		} else {
			if 0x80 <= rune && rune <= 0x9F {
				return false
			}

			if rune < ' ' {
				switch rune {
				case '\n', '\r', '\t':
				default:
					return false
				}
			}

			b = b[size:]
		}
	}
	return true
}

func fileExists(path string) bool {
	var f *os.FileInfo
	var err os.Error

	if f, err = os.Stat(path); err != nil {
		return false
	}

	return f.IsRegular()
}

func directoryExists(path string) bool {
	var f *os.FileInfo
	var err os.Error

	if f, err = os.Stat(path); err != nil {
		return false
	}

	return f.IsDirectory()
}
