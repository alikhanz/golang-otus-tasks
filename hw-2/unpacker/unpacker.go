package unpacker

import (
	"errors"
	"strconv"
	"strings"
)

const EscapeChar = '\\'

type UnPacker struct {

}

func New() UnPacker {
	return UnPacker{}
}

func (u UnPacker) Unpack(s string) (string, error) {
	result := strings.Builder{}

	var previousResultChar rune
	var previousChar rune
	var mustBeEscaped bool

	sRepeatCount := strings.Builder{}
	stringLength := len([]rune(s))

	i := 0
	for _, c := range s {
		if i > 0 {
			previousChar = rune(s[i-1])
		}

		if isEscapeChar(previousChar) && (isNumber(c) || isEscapeChar(c)) && !isEscapeChar(previousResultChar) {
			mustBeEscaped = true
		} else {
			mustBeEscaped = false
		}

		i++
		isLastChar := i == stringLength

		if isLastChar && isEscapeChar(c) && !mustBeEscaped {
			return "", errors.New("некорректная строка")
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

// unicode.IsDigit() не всегда ведет себя так как ожидается, поэтому костылим своё.
// @link https://stackoverflow.com/questions/22593259/check-if-string-is-int-golang#comment75970861_41044460
func isNumber(c rune) bool {
	return c >= 48 && c <= 57
}

func isEscapeChar(c rune) bool {
	return c == EscapeChar
}