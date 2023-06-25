package easy

import (
	"strings"
	"unicode"
)

/*
Фраза является палиндромом, если после преобразования всех прописных букв в строчные
и удаления всех не буквенно-цифровых символов она читается одинаково вперед и назад.
Буквенно-цифровые символы включают буквы и цифры.
Учитывая строку s, вернуть true, если это палиндром, или false в противном случае.
*/

/*
Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Input: "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

func IsPalindrome(s string) bool {

	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	s = strings.Map(f, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
