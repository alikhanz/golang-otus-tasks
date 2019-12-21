package unpacker

import (
	"errors"
	"strconv"
	"strings"
)

type UnPacker struct {
}

func New() UnPacker {
	return UnPacker{}
}

func (u UnPacker) Unpack(s string) (string, error) {
	b := strings.Builder{}

	var lastChar rune

	sRepeatCount := strings.Builder{}
	stringLength := len([]rune(s))

	i := 0
	for _, c := range s {
		i++
		isLastChar := i == stringLength

		// До тех пор пока встречаем число - склеиваем в одну строку, для дальнейшей распаковки.
		if isNumber(c) {

			// Случай когда перед числом нет строкового символа. Возвращаем ошибку.
			if lastChar == 0 {
				return "", errors.New("некорректная строка")
			}

			_, err := sRepeatCount.WriteRune(c)

			if err != nil {
				return "", err
			}

			// Если окажется что последний символ в строке это число, то нужно размножить
			// предыдущий встреченный строковый символ, поэтому продолжаем выполнение.
			if !isLastChar {
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
				b.WriteRune(lastChar)
			}
		}

		if !isNumber(c) {
			b.WriteRune(c)
			lastChar = c
		}
	}

	return b.String(), nil
}

// unicode.IsDigit() не всегда ведет себя так как ожидается, поэтому костылим своё.
// @link https://stackoverflow.com/questions/22593259/check-if-string-is-int-golang#comment75970861_41044460
func isNumber(c rune) bool {
	return c >= 48 && c <= 57
}
