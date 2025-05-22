package main

func letterCombinations(digits string) []string {
	numberMap := make(map[string]string)
	numberMap["1"] = ""
	numberMap["2"] = "abc"
	numberMap["3"] = "def"
	numberMap["4"] = "ghi"
	numberMap["5"] = "jkl"
	numberMap["6"] = "mno"
	numberMap["7"] = "qprs"
	numberMap["8"] = "tuv"
	numberMap["9"] = "wxyz"
	res := make([]string, 0)
	for _, b := range []byte(digits) {
		s := string(b)
		char := numberMap[s]
		if len(res) == 0 {
			for _, c := range []byte(char) {
				res = append(res, string(c))
			}
		} else {
			res = mergeL(res, char)
		}
	}
	return res
}
func mergeL(l []string, char string) []string {
	res := make([]string, 0)
	for _, c := range char {
		for _, s := range l {
			res = append(res, s+string(c))
		}
	}
	return res
}

func main() {
	letterCombinations("23")
}
