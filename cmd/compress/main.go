package main

import (
	"huffman/internal/compress"
)

func main() {
	in := "/home/per/code/huffman/assets/text.txt"
	out := "/home/per/code/huffman/assets/text.scf"

	compress.Compress(in, out)
}
