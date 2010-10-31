package main

import "bytes"

var (
	bNewLine = [2][]byte{{'\n'}, {' '}}
	bSpace   = [2][]byte{{' ', ' '}, {' '}}
	bTag     = [2][]byte{{'>', ' ', '<'}, {'>', '<'}}
)

// Filter html output. This function removes excess whitespace and newlines.
// None of it is required for proper display of the page and only serves to
// bloat the amount of data sent over the wire. This function does not take
// data in <pre>, cdata  and comment tags into account.
func postProcess(data []byte) []byte {
	data = bytes.Replace(data, bNewLine[0], bNewLine[1], -1)

	for bytes.Index(data, bSpace[0]) != -1 {
		data = bytes.Replace(data, bSpace[0], bSpace[1], -1)
	}

	data = bytes.Replace(data, bTag[0], bTag[1], -1)
	return data
}
