package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 3.10
// comma is a version of the comma function using for loops and bytes.Buffer type
// does not anticipate non-integer values, decimals, or sign
func comma(s string) string {
	var buf bytes.Buffer
	length := len(s)
	for i := 0; i < length; i++ {
		// iterate backwards
		index := length - 1 - i
		buf.WriteByte(s[i])

		// every third rune, not first rune in string
		if (index%3) == 0 && (index != 0) {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}

// 3.11
// can deal with proper floating point numbers and signed numbers
func commaWithFloatsAndSigns(s string) string {
	decimalSuffix := ""
	if decimalIndex := strings.Index(s, "."); decimalIndex != -1 {
		decimalSuffix = s[decimalIndex:]
		s = s[:decimalIndex]
	}

	signPrefix := ""
	if s[0] == '-' {
		signPrefix = "-"
		s = s[1:]
	}

	return fmt.Sprintf("%s%s%s", signPrefix, comma(s), decimalSuffix)
}

var testCommaInputs = []string{"1233561", "111", "0", ""}
var testCommaWithFloatsAndSignsInputs = []string{"-1234.567", "1234.55", "0", "-12"}

func main() {
	// 3.10
	fmt.Printf("\nTesting func comma:\n")
	for _, v := range testCommaInputs {
		fmt.Printf("Before: %10s\tAfter: %10s\n", v, comma(v))
	}
	fmt.Println("---")

	// 3.11
	fmt.Println("")
	for _, v := range testCommaWithFloatsAndSignsInputs {
		fmt.Printf("Before: %10s\tAfter: %10s\n", v, commaWithFloatsAndSigns(v))
	}

	return
}
