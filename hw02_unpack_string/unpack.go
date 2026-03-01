package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runes := []rune(str)
	var res strings.Builder

	for i := 0; i < len(runes); i++ {
		if runes[0] >= '0' && runes[0] <= '9' {
			return "", ErrInvalidString
		}

		// Если элемент - цифра, то повторяем предыдущий элемент какое-то кол-во раз и записываем
		// Если два элемента подряд цифры -- ошибка
		if i+1 < len(runes) && runes[i+1] >= '0' && runes[i+1] <= '9' {
			if i+2 < len(runes) && runes[i+2] >= '0' && runes[i+2] <= '9' {
				return "", ErrInvalidString
			}
			num, _ := strconv.Atoi(string(runes[i+1]))
			for j := 0; j < num; j++ {
				res.WriteRune(runes[i])
			}
			i++
			continue
		}
		// Если текущий и предыдущий элементы просто символы
		res.WriteRune(runes[i])
	}
	return res.String(), nil
}
