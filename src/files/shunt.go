//Author: Kieran O'Halloran
//Adapted from: https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
package files

func Intopost(infix string) string {
	specials := map[rune]int{'"': 10, '.': 9, '|': 8}
	var s []rune
	var pofix []rune

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)
		}
	}
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}

// Remove ending of string
func TrimEndString(s string) string {
	if len(s) > 0 {
		s = s[:len(s)-2]
	}
	return s
}
