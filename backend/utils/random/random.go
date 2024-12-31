package random

import (
	crand "crypto/rand"
)

const (
	rs6Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

func SecureAlphaNumeric(n int) string {
	b := make([]byte, n)

	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; {
		idx := int(b[i] & rs6LetterIdxMask)
		if idx < rs6LetterIdxMax {
			b[i] = rs6Letters[idx]
			i++
		} else {
			if _, err := crand.Read(b[i : i+1]); err != nil {
				panic(err)
			}
		}
	}

	return string(b)
}