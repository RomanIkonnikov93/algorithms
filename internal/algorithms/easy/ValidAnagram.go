package easy

import (
	"reflect"
)

/*
Имея две строки s и t, вернуть true, если t является анаграммой s,
и false в противном случае. Анаграмма — это слово или фраза,
образованная путем перестановки букв другого слова или фразы,
обычно с использованием всех исходных букв ровно один раз.
*/

/*
Input: "anagram","nagaram"
Output: true

Input: "rat","car"
Output: false
*/

func IsAnagramMap(s string, t string) bool {

	smap := make(map[rune]int)
	tmap := make(map[rune]int)

	for _, val := range s {
		smap[val]++
	}

	for _, val := range t {
		tmap[val]++
	}

	// or fmt.Sprint(smap) == fmt.Sprint(tmap)
	res := reflect.DeepEqual(smap, tmap)

	return res
}
