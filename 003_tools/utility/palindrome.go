package utility

func isPalindrome(s string) bool {

	r := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		if r[i] != r[len(s)-(i+1)] {
			return false
		}
	}
	return true
}

