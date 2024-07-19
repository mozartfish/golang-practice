package gobyexample

import (
	"fmt"
	"unicode/utf8"
)

// compare a rune value to a rune literal directly
// values enclosed in single quotes are rune literals
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// A Go string is a read-only slice of bytes
// The language and the standard library treat strings specially - as containers of text
// encoded in UTF-8. In other languages, strings are made of characters. In Go the concept
// of a character is called a rune - an integer that represents a unicode code point
func StringsRunes() {
	fmt.Println("STRINGS AND RUNES")
	// s is a string assigned a literal value representing hello in Thai
	// Go string literals are UTF-8 encoded in text
	const s = "สวัสดี"

	// the lenght of the raw bytes stored within s -> strings are equivalent to []byte
	fmt.Println("Len:", len(s))
	
	// indexing into a string produces the raw byte values at each index
	// this loop generates the hex values of all the bytes that constitute code points in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// count how many runes are in a string
	// the runtime of RuneCountInString depends on the size of the string because it has to decode each UTF-8 rune sequentially
	// some Thai characters are represented with multiple UTF-8 code points that can span multiple bytes
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range loops handle strings specially and decode each rune along with its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")

	// same things as range loop but using DecodeRuneInString function
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// pass rune value to a function
		examineRune(runeValue)
	}
}
