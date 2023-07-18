package medium

import "fmt"

/*
Учитывая массив строк strs, сгруппируйте анаграммы вместе. Вы можете вернуть ответ в любом порядке.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы,
обычно с использованием всех исходных букв ровно один раз.
*/

/*
Input: []string{"eat","tea","tan","ate","nat","bat"}
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Input: []string{""}
Output: [[""]]

Input: []string{"a"}
Output: [["a"]]
*/

func GroupAnagrams(strs []string) [][]string {

	res := make([][]string, 0)

	if len(strs) == 1 {
		res = append(res, strs)
		return res
	}

	s := make(map[int]string)
	w := make(map[int]string)

	for i, v := range strs {
		f := make(map[string]int)
		for _, r := range v {
			f[string(r)]++
		}
		word := fmt.Sprint(f)
		w[i] = word
		s[i] = v
	}

	for k := range s {
		arr := make([]string, 0)
		for i := len(w); i >= 0; i-- {
			if _, ok := s[k]; ok {
				arr = append(arr, s[k])
				delete(s, k)
			}
			if w[k] == w[i] {
				if _, ok := s[i]; ok {
					arr = append(arr, s[i])
					delete(s, i)
				}
			}
		}
		if len(arr) > 0 {
			res = append(res, arr)
		}
	}

	return res
}
