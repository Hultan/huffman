package main

import (
	"compress/internal/compress"
)

func main() {
	in := "/home/per/code/compress/assets/text.txt"
	out := "/home/per/code/compress/assets/text.scf"

	compress.Compress(in, out)
}
