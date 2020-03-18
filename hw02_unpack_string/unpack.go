package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder
	runes := []rune(s)
	var count int
	isEscaped := false

	for i := 0; i < len(runes); i++ {
		count = 1

		switch {
		case runes[i] == '\\' && !isEscaped && i != len(runes)-1:
			isEscaped = true
			continue
		case i == len(runes)-1:
			if runes[i] == '\\' && !isEscaped {
				return "", ErrInvalidString
			}
			if unicode.IsDigit(runes[i]) && !isEscaped {
				continue
			}
		case unicode.IsDigit(runes[i+1]):
			if unicode.IsDigit(runes[i]) && !isEscaped {
				return "", ErrInvalidString
			}
			count, _ = strconv.Atoi(string(runes[i+1]))
		case unicode.IsDigit(runes[i]) && !isEscaped:
			continue
		case unicode.IsDigit(runes[0]):
			return "", ErrInvalidString
		}

		isEscaped = false

		builder.WriteString(strings.Repeat(string(runes[i]), count))
	}

	return builder.String(), nil
}
