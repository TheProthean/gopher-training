package phoneNormalizer

import (
	"errors"
	"strings"

	"github.com/gopher-training/phone_normalizer/setup"
)

//NormalizeNumbers is a function for normalizing numbers and checking duplicates. Instead of duplicates you will find empty strings
func NormalizeNumbers(numbers []setup.PhoneNumber) ([]setup.PhoneNumber, error) {
	foundNumbers := map[string]bool{}
	for i := range numbers {
		if numbers[i].PhoneNumber == "" {
			return nil, errors.New("There was an empty field instead of phone number")
		}
		mappedNumber := strings.Map(func(r rune) rune {
			switch {
			case r >= '0' && r <= '9':
				return r
			default:
				return 'a'
			}
		}, numbers[i].PhoneNumber)
		normalizedNumber := strings.Replace(mappedNumber, "a", "", -1)
		if _, ok := foundNumbers[normalizedNumber]; !ok {
			numbers[i].PhoneNumber = normalizedNumber
			foundNumbers[normalizedNumber] = true
		} else {
			numbers[i].PhoneNumber = ""
		}
	}
	return numbers, nil
}
