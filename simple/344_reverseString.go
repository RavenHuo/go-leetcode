package main

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	i := 0
	j := len(s) - 1
	for {
		if i >= j {
			break
		}
		c := s[i]
		s[i] = s[j]
		s[j] = c
		i++
		j--
	}
}

func main() {
	reverseString([]byte("hello"))
}
