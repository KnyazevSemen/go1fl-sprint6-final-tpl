package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertTextOrMorse(input string) (string, error) {
	s := strings.TrimSpace(input)

	if s == "" {
		return "", errors.New("передана пустая строка")
	}

	isMorse := isMorse(s)

	if isMorse {
		result := morse.ToText(s)
		return result, nil
	} else {
		result := morse.ToMorse(s)
		return result, nil
	}
}

func isMorse(s string) bool {
	for _, char := range s {
		if char != '.' && char != '-' && char != ' ' && char != '/' {
			return false
		}
	}

	return true

}
