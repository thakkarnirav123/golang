package utility

func isPalindrome(s string) bool {
	result := true
	if s == "" {
		result = false
	}else{
		r := []rune(s)
		for i := 0; i < len(s)/2; i++ {
			if r[i] != r[len(s)-(i+1)] {
				result = false
				break
			}
		}
	}

	return result
}

