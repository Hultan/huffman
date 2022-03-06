package compress

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Pair struct {
	Key   byte
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func Compress(in, out string) {
	data, err := getDataToCompress(in)
	if err != nil {
		panic(err)
	}

	// Create a map of rune frequencies
	freq := getFrequencyMap(data)
	// Get encoding map
	encodingMap := getEncodingMap(freq)

	// Print the encoding map
	printEncodingMap(encodingMap)

	compressedData := compressData(string(data), encodingMap)

	fmt.Println(compressedData)
	fmt.Println(len(string(data)), len(compressedData))

	// f, _ := os.Open(in)
	// data, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
}

func getEncodingMap(freq map[rune]int) map[rune]string {
	// Build a Huffman tree from frequency map
	tree := buildHuffmanTree(freq)
	// Create a encoding map
	encodingMap := make(map[rune]string)
	// Get the encoding map
	getEncodingMapFromTree(tree, []byte{}, encodingMap)

	return encodingMap
}

func getDataToCompress(in string) ([]byte, error) {
	f, _ := os.Open(in)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	f.Close()

	return data, nil
}

func compressData(data string, codeMap map[rune]string) string {
	result := ""
	for _, c := range data {
		result += codeMap[c] + " "
	}
	return result
}

func printEncodingMap(codeMap map[rune]string) {
	for r, s := range codeMap {
		fmt.Println(string(r), s)
	}
}

func getFrequencyMap(data []byte) map[rune]int {
	freq := make(map[rune]int)

	for _, c := range string(data) {
		freq[c]++
	}

	return freq
}
