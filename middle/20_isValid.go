/**
 * @Author raven
 * @Description
 * @Date 2022/6/13
 **/
package main

func isValid(s string) bool {
	charArr := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		sChar := s[i : i+1]
		if sChar == "(" || sChar == "{" || sChar == "[" {
			charArr = append(charArr, sChar)
		} else {
			if i == 0 {
				return false
			}

			if len(charArr) <= 0 {
				return false
			}
			first := charArr[len(charArr)-1:][0]
			if !match(sChar, first) {
				return false
			} else {
				charArr = charArr[:len(charArr)-1]
			}
		}
	}
	return len(charArr) == 0
}

func match(second, first string) bool {
	if first == "(" && second == ")" {
		return true
	}
	if first == "{" && second == "}" {
		return true
	}
	if first == "[" && second == "]" {
		return true
	}
	return false
}

func main() {
	isValid("{[]}")
}
