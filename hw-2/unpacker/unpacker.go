package unpacker

import (
	"errors"
	"strconv"
	"strings"
)

// EscapeChar символ экранирования.
const EscapeChar = '\\'

// UnPacker структура распаковщика.
type UnPacker struct {

}

// New функция конструктор.
func New() UnPacker {
	return UnPacker{}
}

// Unpack функция распаковки строки.
func (u UnPacker) Unpack(s string) (string, error) {
	result := strings.Builder{}

	var previousResultChar rune
	var previousChar rune
	var mustBeEscaped bool
	sRune := []rune(s)

	sRepeatCount := strings.Builder{}
	stringLength := len(sRune)

	for i, c := range sRune {
		if i > 0 {
			previousChar = sRune[i-1]
		}

		if isEscapeChar(previousChar) && !mustBeEscaped {
			mustBeEscaped = true
		} else {
			mustBeEscaped = false
		}

		if mustBeEscaped && !canBeEscaped(c) {
			return "", errors.New("экранирован некорректный символ")
		}

		isLastChar := i == stringLength-1

		if isLastChar && isEscapeChar(c) && !mustBeEscaped {
			return "", errors.New("некорректная строка")
		}

		if isEscapeChar(c) && !mustBeEscaped {
			continue
		}

		if isNumber(c) && !mustBeEscaped {
			// До тех пор пока встречаем неэкранированное число - склеиваем в одну строку, для дальнейшей распаковки.

			// Случай когда число идет первым символом в строке. Возвращаем ошибку.
			if previousResultChar == 0 {
				return "", errors.New("некорректная строка")
			}

			_, err := sRepeatCount.WriteRune(c)

			if err != nil {
				return "", err
			}

			if !isLastChar {
				// В случае когда это последний символ в строке - продолжаем выполнение цикла для распаковки
				// последнего встреченного символа
				continue
			}
		}

		if sRepeatCount.Len() > 0 {
			repeatCount, err := strconv.Atoi(sRepeatCount.String())
			sRepeatCount.Reset()

			if err != nil {
				return "", err
			}

			for j := 0; j < repeatCount-1; j++ {
				result.WriteRune(previousResultChar)
			}
		}

		if (!isNumber(c) && !isEscapeChar(c)) || mustBeEscaped {
			result.WriteRune(c)
			previousResultChar = c
		}
	}
	return result.String(), nil
}

func canBeEscaped(c rune) bool {
	return isNumber(c) || isEscapeChar(c)
}

// unicode.IsDigit() не всегда ведет себя так как ожидается, поэтому костылим своё.
// @link https://stackoverflow.com/questions/22593259/check-if-string-is-int-golang#comment75970861_41044460
func isNumber(c rune) bool {
	return c >= 48 && c <= 57
}

func isEscapeChar(c rune) bool {
	return c == EscapeChar
}