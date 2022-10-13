// Based on Element Android source
// https://github.com/vector-im/element-android/blob/02a8fd231c1e517a7cc6af1ea7127a0595ba64fd/vector/src/main/java/im/vector/app/features/home/room/detail/timeline/helper/MatrixItemColorProvider.kt

package main

import (
	"fmt"
	"os"
)

const (
	CLR_1 = "#368BD6" // Azure
	CLR_2 = "#AC3BA8" // Grape
	CLR_3 = "#03B381" // Verde
	CLR_4 = "#E64F7A" // Polly
	CLR_5 = "#FF812D" // Melon
	CLR_6 = "#2DC2C5" // Aqua
	CLR_7 = "#5C56F5" // Prune
	CLR_8 = "#74D12C" // Kiwi
)

func GetMxidHash(mxid string) int32 {
	var hash int32
	mxidChars := []rune(mxid)

	for _, chr := range mxidChars {
		hash = (hash << 5) - hash + int32(chr)
	}

	return hash
}

func GetColorNumberFromHash(hash int32) int {
	if hash < 0 {
		hash *= -1
	}

	return int(hash % 8)
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: %s <mxid>\n", args[0])
		return
	}

	fmt.Printf("MXID:            %s\n", args[1])

	hash := GetMxidHash(args[1])

	fmt.Printf("MXID HASH:       %d\n", hash)

	colorNumber := GetColorNumberFromHash(hash)

	fmt.Printf("COLOR NUMBER:    %d\n", colorNumber+1)

	fmt.Printf("HEX COLOR CODE:  %s\n", map[int]string{
		0: CLR_1,
		1: CLR_2,
		2: CLR_3,
		3: CLR_4,
		4: CLR_5,
		5: CLR_6,
		6: CLR_7,
		7: CLR_8,
	}[colorNumber])
}
